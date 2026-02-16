package router

import (
	"smart-campus-services/handlers"

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

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandlers.Register)
		auth.POST("/login", authHandlers.Login)
		auth.POST("/logout", authHandlers.Logout)
		auth.POST("/refresh", authHandlers.RefreshToken)
	}

	users := r.Group("/api/users")
	{
		users.GET("/:id", userHandlers.GetUser)
		users.PUT("/:id", userHandlers.UpdateUser)
		users.GET("/:id/profile", userHandlers.GetProfile)
	}

	services := r.Group("/api/services")
	{
		services.GET("", serviceHandlers.ListServices)
		services.GET("/:id", serviceHandlers.GetService)
		services.POST("", serviceHandlers.CreateService)
		services.PUT("/:id", serviceHandlers.UpdateService)
		services.DELETE("/:id", serviceHandlers.DeleteService)
		services.GET("/category/:category", serviceHandlers.GetServicesByCategory)
	}

	bookings := r.Group("/api/bookings")
	{
		bookings.POST("", bookingHandlers.CreateBooking)
		bookings.GET("/:id", bookingHandlers.GetBooking)
		bookings.GET("/user/:userId", bookingHandlers.GetUserBookings)
		bookings.PUT("/:id", bookingHandlers.UpdateBooking)
		bookings.DELETE("/:id", bookingHandlers.CancelBooking)
	}

	approval := r.Group("/api/approval")
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
	{
		notifications.GET("/:userId", notificationHandlers.GetNotifications)
		notifications.POST("", notificationHandlers.CreateNotification)
		notifications.PUT("/:id/read", notificationHandlers.MarkAsRead)
	}

	reviews := r.Group("/api/reviews")
	{
		reviews.POST("", reviewHandlers.CreateReview)
		reviews.GET("/service/:serviceId", reviewHandlers.GetServiceReviews)
		reviews.GET("/:id", reviewHandlers.GetReview)
		reviews.DELETE("/:id", reviewHandlers.DeleteReview)
	}
}
