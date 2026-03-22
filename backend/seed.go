package main

import (
	"smart-campus-services/models"

	"gorm.io/gorm"
)

// seedDatabase adds a small, idempotent sample dataset for local development.
func seedDatabase(db *gorm.DB) error {
	var serviceCount int64
	if err := db.Model(&models.Service{}).Count(&serviceCount).Error; err != nil {
		return err
	}

	if serviceCount > 0 {
		return nil
	}

	service := models.Service{
		Name:        "Campus Library",
		Description: "General library support and study-space assistance.",
		Category:    "library",
		Location:    "Main Library",
		Email:       "library@campus.edu",
		Hours:       "Mon-Fri 8:00 AM - 8:00 PM",
		IsActive:    true,
	}
	if err := db.Create(&service).Error; err != nil {
		return err
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
			Email:     "admin@campus.edu",
			Password:  "password123",
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
