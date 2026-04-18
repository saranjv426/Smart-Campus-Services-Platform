package models

import "testing"

func TestUserBeforeCreateAssignsID(t *testing.T) {
	user := User{}

	if err := user.BeforeCreate(nil); err != nil {
		t.Fatalf("BeforeCreate returned error: %v", err)
	}

	if user.ID == "" {
		t.Fatal("expected user ID to be assigned")
	}
}

func TestServiceBeforeCreateAssignsID(t *testing.T) {
	service := Service{}

	if err := service.BeforeCreate(nil); err != nil {
		t.Fatalf("BeforeCreate returned error: %v", err)
	}

	if service.ID == "" {
		t.Fatal("expected service ID to be assigned")
	}
}

func TestBookingBeforeCreateAssignsID(t *testing.T) {
	booking := Booking{}

	if err := booking.BeforeCreate(nil); err != nil {
		t.Fatalf("BeforeCreate returned error: %v", err)
	}

	if booking.ID == "" {
		t.Fatal("expected booking ID to be assigned")
	}
}

func TestNotificationBeforeCreateAssignsID(t *testing.T) {
	notification := Notification{}

	if err := notification.BeforeCreate(nil); err != nil {
		t.Fatalf("BeforeCreate returned error: %v", err)
	}

	if notification.ID == "" {
		t.Fatal("expected notification ID to be assigned")
	}
}

func TestReviewBeforeCreateAssignsID(t *testing.T) {
	review := Review{}

	if err := review.BeforeCreate(nil); err != nil {
		t.Fatalf("BeforeCreate returned error: %v", err)
	}

	if review.ID == "" {
		t.Fatal("expected review ID to be assigned")
	}
}
