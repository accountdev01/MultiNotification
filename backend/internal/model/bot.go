package model

import "github.com/google/uuid"

type Bot struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()" db:"id"`
	Name      string    `json:"name" gorm:"size:50;not null" db:"name" binding:"required"`
	Platform  string    `json:"platform" gorm:"size:20;not null" db:"platform" binding:"required"`
	Token     string    `json:"token" gorm:"size:255;not null" db:"token" binding:"required"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true" db:"is_active"`
	TotalSent int64     `json:"total_counter" gorm:"not null;default:0" db:"total_sent_count"`
	RateLimit int64     `json:"rate_limit" db:"rate_limit"`
}
