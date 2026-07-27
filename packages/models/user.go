package models

import (
	"time"

	"gorm.io/gorm"
)

// Status represents the state of the user's account verification
type Status string

const (
	StatusUnverified Status = "UNVERIFIED"
	StatusVerified   Status = "VERIFIED"
)

type User struct {
	ID          string    `gorm:"primaryKey;type:varchar(128)" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Surname     string    `gorm:"type:varchar(100);not null" json:"surname"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	DateOfBirth time.Time `gorm:"type:date;not null" json:"dateOfBirth"`

	// Foreign Key / Relation reference
	VerificationStatusID *string             `gorm:"type:uuid" json:"verificationStatusId,omitempty"`
	VerificationStatus   *VerificationStatus `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"verificationStatus,omitempty"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type VerificationStatus struct {
	ID               string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	VerificationCode string    `gorm:"type:varchar(255);not null" json:"-"` // Bcrypt hash of code (hidden in JSON)
	Status           Status    `gorm:"type:varchar(20);default:'UNVERIFIED';not null" json:"status"`
	UserID           string    `gorm:"type:uuid;uniqueIndex;not null" json:"userId"`
	ExpiresAt        time.Time `gorm:"not null" json:"expiresAt"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
