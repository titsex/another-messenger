package models

import (
	"time"
)

type User struct {
	ID             string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	CreatedAt      time.Time  `json:"created_at" example:"2023-01-01T12:00:00Z"`
	UpdatedAt      time.Time  `json:"updated_at" example:"2023-01-01T12:00:00Z"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" example:"2023-01-01T12:00:00Z"`
	LastSeenAt     *time.Time `json:"last_seen_at,omitempty" example:"2023-01-01T12:00:00Z"`
	Name           string     `json:"name" example:"John Doe"`
	Email          string     `json:"email" example:"john.doe@example.com"`
	Password       string     `json:"-"`
	AvatarURL      string     `json:"avatar_url,omitempty" example:"https://example.com/avatar.jpg"`
	EmailConfirmed bool       `json:"email_confirmed" example:"true"`
	Username       string     `json:"username" example:"johndoe"`
	Bio            string     `json:"bio,omitempty" example:"Software developer"`
	PhoneNumber    string     `json:"phone_number,omitempty" example:"+1234567890"`
}
