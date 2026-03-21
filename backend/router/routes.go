package router

import (
	"smart-campus-services/handlers"
	"smart-campus-services/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAPIRoutes attaches all API route groups to the given engine.
func RegisterAPIRoutes(r *gin.Engine, db *gorm.DB) {
	authHandlers := handlers.NewAuthHandler(db)
	serviceHandlers := handlers.NewServiceHandler(db)
	bookingHandlers := handlers.NewBookingHandler(db)
	notificationHandlers := handlers.NewNotificationHandler(db)
	reviewHandlers := handlers.NewReviewHandler(db)
	userHandlers := handlers.NewUserHandler(db)
	approvalHandlers := handlers.NewApprovalHandler(db)
	authRequired := middleware.AuthRequired()

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandlers.Register)
		auth.POST("/login", authHandlers.Login)
	}
	authProtected := r.Group("/api/auth")
	authProtected.Use(authRequired)
	{
		authProtected.POST("/logout", authHandlers.Logout)
		authProtected.POST("/refresh", authHandlers.RefreshToken)
	}

	users := r.Group("/api/users")
	{
		users.GET("/:id", userHandlers.GetUser)
		users.GET("/:id/profile", userHandlers.GetProfile)
	}
	usersProtected := r.Group("/api/users")
	usersProtected.Use(authRequired)
	{
		usersProtected.PUT("/:id", userHandlers.UpdateUser)
	}

	services := r.Group("/api/services")
	{
		services.GET("", serviceHandlers.ListServices)
		services.GET("/:id", serviceHandlers.GetService)
		services.GET("/category/:category", serviceHandlers.GetServicesByCategory)
	}
	servicesProtected := r.Group("/api/services")
	servicesProtected.Use(authRequired)
	{
		servicesProtected.POST("", serviceHandlers.CreateService)
		servicesProtected.PUT("/:id", serviceHandlers.UpdateService)
		servicesProtected.DELETE("/:id", serviceHandlers.DeleteService)
	}

	bookings := r.Group("/api/bookings")
	{
		bookings.GET("/:id", bookingHandlers.GetBooking)
		bookings.GET("/user/:userId", bookingHandlers.GetUserBookings)
	}
	bookingsProtected := r.Group("/api/bookings")
	bookingsProtected.Use(authRequired)
	{
		bookingsProtected.POST("", bookingHandlers.CreateBooking)
		bookingsProtected.PUT("/:id", bookingHandlers.UpdateBooking)
		bookingsProtected.PATCH("/:id/status", bookingHandlers.CancelBooking)
	}

	approval := r.Group("/api/approval")
	approval.Use(authRequired)
	{
		approval.GET("/staff/:staffId/pending", approvalHandlers.GetPendingBookings)
		approval.GET("/staff/:staffId/all", approvalHandlers.GetAllBookingsForService)
		approval.PUT("/bookings/:id/approve", approvalHandlers.ApproveBooking)
		approval.PUT("/bookings/:id/reject", approvalHandlers.RejectBooking)
		approval.GET("/admin/:userId/pending", approvalHandlers.GetAllPendingBookings)
		approval.GET("/admin/:userId/all", approvalHandlers.GetAllBookings)
		approval.PUT("/admin/:userId/bookings/:id/approve", approvalHandlers.AdminApproveBooking)
		approval.PUT("/admin/:userId/bookings/:id/reject", approvalHandlers.AdminRejectBooking)
	}

	notifications := r.Group("/api/notifications")
	notifications.Use(authRequired)
	{
		notifications.GET("/:userId", notificationHandlers.GetNotifications)
		notifications.POST("", notificationHandlers.CreateNotification)
		notifications.PUT("/:id/read", notificationHandlers.MarkAsRead)
	}

	reviews := r.Group("/api/reviews")
	{
		reviews.GET("/service/:serviceId", reviewHandlers.GetServiceReviews)
		reviews.GET("/:id", reviewHandlers.GetReview)
	}
	reviewsProtected := r.Group("/api/reviews")
	reviewsProtected.Use(authRequired)
	{
		reviewsProtected.POST("", reviewHandlers.CreateReview)
		reviewsProtected.DELETE("/:id", reviewHandlers.DeleteReview)
	}
}
