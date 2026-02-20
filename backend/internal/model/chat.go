package model

import "github.com/google/uuid"

type Chat struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()" db:"id"`
	Name      string    `json:"name" gorm:"size:50;not null" db:"name"`
	Token     string    `json:"token" gorm:"size:255;not null" db:"token"`
	Bot       string    `json:"bot" gorm:"size:255;not null" db:"bot"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true" db:"is_active"`
	TotalSent int64     `json:"total_sent" gorm:"not null;default:0" db:"total_sent_count"`
	RateLimit int64     `json:"rate_limit" db:"rate_limit"`
}
