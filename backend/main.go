package main

import (
	"log"
	"os"
	"path/filepath"

	"smart-campus-services/middleware"
	"smart-campus-services/models"
	"smart-campus-services/router"
	"smart-campus-services/validation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	if ginMode := os.Getenv("GIN_MODE"); ginMode != "" {
		gin.SetMode(ginMode)
	}

	if err := validation.Init(); err != nil {
		log.Fatalf("Failed to initialize validator: %v", err)
	}

	db, err := initializeDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.Booking{},
		&models.Notification{},
		&models.Review{},
	); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	r := gin.New()
	r.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.CORS(),
		middleware.RequestLogger(),
		middleware.ErrorHandler(),
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"message": "Smart Campus Services Platform is running",
		})
	})

	r.POST("/api/seed", func(c *gin.Context) {
		if err := seedDatabase(db); err != nil {
			c.JSON(500, gin.H{"error": "Failed to seed database"})
			return
		}
		c.JSON(200, gin.H{"message": "Database seeded successfully"})
	})

	router.RegisterAPIRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
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
