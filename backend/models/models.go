package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a campus user
type User struct {
	ID         string    `gorm:"primaryKey" json:"id"`
	Email      string    `gorm:"uniqueIndex" json:"email"`
	Password   string    `json:"-"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Phone      string    `json:"phone"`
	Role       string    `json:"role"` // student, staff, admin
	Department string    `json:"department"`
	AvatarURL  string    `json:"avatarUrl"`
	Bio        string    `json:"bio"`
	ServiceID  string    `json:"serviceId"` // If staff, which service they manage
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`

	Bookings      []Booking      `gorm:"foreignKey:UserID" json:"bookings,omitempty"`
	Reviews       []Review       `gorm:"foreignKey:UserID" json:"reviews,omitempty"`
	Notifications []Notification `gorm:"foreignKey:UserID" json:"notifications,omitempty"`
}

// BeforeCreate hook for User
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}

// Service represents a campus service
type Service struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"` // library, dining, transportation, maintenance, etc.
	ImageURL    string    `json:"imageUrl"`
	Location    string    `json:"location"`
	Phone       string    `json:"phone"`
	Email       string    `json:"email"`
	Hours       string    `json:"hours"`
	Rating      float64   `json:"rating"`
	IsActive    bool      `json:"isActive"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

	Bookings []Booking `gorm:"foreignKey:ServiceID" json:"bookings,omitempty"`
	Reviews  []Review  `gorm:"foreignKey:ServiceID" json:"reviews,omitempty"`
}

// BeforeCreate hook for Service
func (s *Service) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// Booking represents a service booking
type Booking struct {
	ID            string    `gorm:"primaryKey" json:"id"`
	UserID        string    `json:"userId"`
	ServiceID     string    `json:"serviceId"`
	Status        string    `json:"status"` // pending, approved, rejected, completed, cancelled
	StartTime     time.Time `json:"startTime"`
	EndTime       time.Time `json:"endTime"`
	Notes         string    `json:"notes"`
	ApprovedBy    string    `json:"approvedBy"`    // Staff ID who approved
	ApprovalNotes string    `json:"approvalNotes"` // Staff notes for approval/rejection
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`

	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Service Service `gorm:"foreignKey:ServiceID" json:"service,omitempty"`
}

// BeforeCreate hook for Booking
func (b *Booking) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return nil
}

// Notification represents a user notification
type Notification struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"` // booking, reminder, announcement, etc.
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`

	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// BeforeCreate hook for Notification
func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == "" {
		n.ID = uuid.New().String()
	}
	return nil
}

// Review represents a service review
type Review struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"userId"`
	ServiceID string    `json:"serviceId"`
	Rating    int       `json:"rating"` // 1-5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User    User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Service Service `gorm:"foreignKey:ServiceID" json:"service,omitempty"`
}

// BeforeCreate hook for Review
func (r *Review) BeforeCreate(tx *gorm.DB) error {
	if r.ID == "" {
		r.ID = uuid.New().String()
	}
	return nil
}
