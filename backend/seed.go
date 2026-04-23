package main

import (
	"smart-campus-services/models"

	"gorm.io/gorm"
)

// seedDatabase adds a small, idempotent sample dataset for local development.
func seedDatabase(db *gorm.DB) error {
	// Delete all existing services to force fresh seeding with image URLs
	if err := db.Exec("DELETE FROM services").Error; err != nil {
		return err
	}
	
	// Create the first service with image URL
	service := models.Service{
		Name:        "Campus Library",
		Description: "General library support and study-space assistance.",
		Category:    "library",
		Location:    "Main Library",
		Email:       "library@campus.edu",
		Hours:       "Mon-Fri 8:00 AM - 8:00 PM",
		IsActive:    true,
		ImageURL:    "https://images.unsplash.com/photo-150784272343-583f20270319?w=300&h=200&fit=crop",
	}
	if err := db.Create(&service).Error; err != nil {
		return err
	}

	// Add more services across different categories
	additionalServices := []models.Service{
		{
			Name:        "Main Dining Hall",
			Description: "Campus dining facility with variety of meals and snacks.",
			Category:    "dining",
			Location:    "Student Center",
			Email:       "dining@campus.edu",
			Phone:       "555-1001",
			Hours:       "Mon-Sun 7:00 AM - 10:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-1567521464027-f127ff144326?w=300&h=200&fit=crop",
		},
		{
			Name:        "Campus Shuttle Service",
			Description: "Free shuttle bus service connecting all campus buildings and nearby stations.",
			Category:    "transportation",
			Location:    "Transportation Hub",
			Email:       "transport@campus.edu",
			Phone:       "555-1002",
			Hours:       "Mon-Fri 6:00 AM - 11:00 PM, Sat-Sun 8:00 AM - 9:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-1464207687429-7505649dae38?w=300&h=200&fit=crop",
		},
		{
			Name:        "Campus Health Center",
			Description: "Medical and health services for students and staff.",
			Category:    "health",
			Location:    "Health Services Building",
			Email:       "health@campus.edu",
			Phone:       "555-1003",
			Hours:       "Mon-Fri 9:00 AM - 6:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-1576091160550-2173dba999ef?w=300&h=200&fit=crop",
		},
		{
			Name:        "Reference Desk",
			Description: "Librarian assistance for research and information requests.",
			Category:    "library",
			Location:    "Main Library - 2nd Floor",
			Email:       "reference@campus.edu",
			Phone:       "555-1004",
			Hours:       "Mon-Fri 10:00 AM - 6:00 PM, Sat 12:00 PM - 4:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-150784272343-583f20270319?w=300&h=200&fit=crop",
		},
		{
			Name:        "Campus Cafe",
			Description: "Quick service cafe with coffee, pastries, and light meals.",
			Category:    "dining",
			Location:    "Multiple Locations",
			Email:       "cafe@campus.edu",
			Phone:       "555-1005",
			Hours:       "Mon-Fri 7:30 AM - 5:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-1559056199-641a0ac8b8d5?w=300&h=200&fit=crop",
		},
		{
			Name:        "Parking Services",
			Description: "Campus parking permit management and vehicle registration.",
			Category:    "transportation",
			Location:    "Transportation Hub",
			Email:       "parking@campus.edu",
			Phone:       "555-1006",
			Hours:       "Mon-Fri 8:00 AM - 5:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-1506521295926-19bfd768e4ef?w=300&h=200&fit=crop",
		},
		{
			Name:        "Counseling Services",
			Description: "Mental health and counseling support for students.",
			Category:    "health",
			Location:    "Student Services Building",
			Email:       "counseling@campus.edu",
			Phone:       "555-1007",
			Hours:       "Mon-Fri 9:00 AM - 5:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-1544716278-ca5e3af4abd8?w=300&h=200&fit=crop",
		},
		{
			Name:        "Study Rooms",
			Description: "Group study and collaboration spaces available for reservation.",
			Category:    "library",
			Location:    "Main Library - Various Floors",
			Email:       "rooms@campus.edu",
			Phone:       "555-1008",
			Hours:       "Mon-Fri 8:00 AM - 8:00 PM, Weekends 10:00 AM - 6:00 PM",
			IsActive:    true,
			ImageURL:    "https://images.unsplash.com/photo-1522202176988-696596bce6ff?w=300&h=200&fit=crop",
		},
	}

	for _, svc := range additionalServices {
		if err := db.Create(&svc).Error; err != nil {
			return err
		}
	}

	users := []models.User{
		{
			Email:     "student@campus.edu",
			Password:  "password123",
			FirstName: "Student",
			LastName:  "User",
			Phone:     "555-0100",
			Role:      "student",
		},
		{
			Email:     "staff@campus.edu",
			Password:  "password123",
			FirstName: "Staff",
			LastName:  "Manager",
			Phone:     "555-0101",
			Role:      "staff",
			ServiceID: service.ID,
		},
		{
			Email:     "admin@ufl.edu",
			Password:  "admin123",
			FirstName: "Admin",
			LastName:  "User",
			Phone:     "555-0102",
			Role:      "admin",
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
