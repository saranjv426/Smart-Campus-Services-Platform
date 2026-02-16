package testutil

import (
	"fmt"
	"testing"

	"smart-campus-services/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewTestDB creates an isolated in-memory SQLite DB for tests.
func NewTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Service{},
		&models.Booking{},
		&models.Notification{},
		&models.Review{},
	); err != nil {
		t.Fatalf("failed to migrate test db: %v", err)
	}

	return db
}
