package testutil

import (
	"testing"

	"smart-campus-services/models"
)

func TestNewTestDBMigratesModels(t *testing.T) {
	db := NewTestDB(t)

	service := models.Service{
		Name:        "Library",
		Description: "Study support",
		Category:    "library",
		Location:    "Main Library",
		IsActive:    true,
	}

	if err := db.Create(&service).Error; err != nil {
		t.Fatalf("expected migrated test database to persist models, got error: %v", err)
	}

	if service.ID == "" {
		t.Fatal("expected model hooks to run in test database")
	}
}
