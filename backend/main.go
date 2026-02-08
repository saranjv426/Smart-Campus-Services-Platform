package main

import (
	"log"
	"os"
	"path/filepath"

	"smart-campus-services/handlers"
	"smart-campus-services/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := db.AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.Booking{},
		&models.Notification{},
		&models.Review{},
	); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(corsMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Smart Campus Services Platform is running",
		})
	})

	// Seed database endpoint (for development only)
	router.POST("/api/seed", func(c *gin.Context) {
		if err := seedDatabase(db); err != nil {
			c.JSON(500, gin.H{"error": "Failed to seed database"})
			return
		}
		c.JSON(200, gin.H{"message": "Database seeded successfully"})
	})

	// Initialize handlers with database connection
	authHandlers := handlers.NewAuthHandler(db)
	serviceHandlers := handlers.NewServiceHandler(db)
	bookingHandlers := handlers.NewBookingHandler(db)
	notificationHandlers := handlers.NewNotificationHandler(db)
	reviewHandlers := handlers.NewReviewHandler(db)
	userHandlers := handlers.NewUserHandler(db)
	approvalHandlers := handlers.NewApprovalHandler(db)

	// Auth routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authHandlers.Register)
		auth.POST("/login", authHandlers.Login)
		auth.POST("/logout", authHandlers.Logout)
		auth.POST("/refresh", authHandlers.RefreshToken)
	}

	// User routes
	users := router.Group("/api/users")
	{
		users.GET("/:id", userHandlers.GetUser)
		users.PUT("/:id", userHandlers.UpdateUser)
		users.GET("/:id/profile", userHandlers.GetProfile)
	}

	// Service routes
	services := router.Group("/api/services")
	{
		services.GET("", serviceHandlers.ListServices)
		services.GET("/:id", serviceHandlers.GetService)
		services.POST("", serviceHandlers.CreateService)
		services.PUT("/:id", serviceHandlers.UpdateService)
		services.DELETE("/:id", serviceHandlers.DeleteService)
		services.GET("/category/:category", serviceHandlers.GetServicesByCategory)
	}

	// Booking routes
	bookings := router.Group("/api/bookings")
	{
		bookings.POST("", bookingHandlers.CreateBooking)
		bookings.GET("/:id", bookingHandlers.GetBooking)
		bookings.GET("/user/:userId", bookingHandlers.GetUserBookings)
		bookings.PUT("/:id", bookingHandlers.UpdateBooking)
		bookings.DELETE("/:id", bookingHandlers.CancelBooking)
	}

	// Approval routes (staff only)
	approval := router.Group("/api/approval")
	{
		approval.GET("/staff/:staffId/pending", approvalHandlers.GetPendingBookings)
		approval.GET("/staff/:staffId/all", approvalHandlers.GetAllBookingsForService)
		approval.PUT("/bookings/:id/approve", approvalHandlers.ApproveBooking)
		approval.PUT("/bookings/:id/reject", approvalHandlers.RejectBooking)
		// Admin approval routes
		approval.GET("/admin/:userId/pending", approvalHandlers.GetAllPendingBookings)
		approval.GET("/admin/:userId/all", approvalHandlers.GetAllBookings)
		approval.PUT("/admin/:userId/bookings/:id/approve", approvalHandlers.AdminApproveBooking)
		approval.PUT("/admin/:userId/bookings/:id/reject", approvalHandlers.AdminRejectBooking)
	}

	// Notification routes
	notifications := router.Group("/api/notifications")
	{
		notifications.GET("/:userId", notificationHandlers.GetNotifications)
		notifications.POST("", notificationHandlers.CreateNotification)
		notifications.PUT("/:id/read", notificationHandlers.MarkAsRead)
	}

	// Review routes
	reviews := router.Group("/api/reviews")
	{
		reviews.POST("", reviewHandlers.CreateReview)
		reviews.GET("/service/:serviceId", reviewHandlers.GetServiceReviews)
		reviews.GET("/:id", reviewHandlers.GetReview)
		reviews.DELETE("/:id", reviewHandlers.DeleteReview)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func initializeDatabase() (*gorm.DB, error) {
	dbPath := databasePath()
	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databasePath() string {
	if dbPath := os.Getenv("DB_PATH"); dbPath != "" {
		return dbPath
	}

	return "data/smart_campus.db"
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
