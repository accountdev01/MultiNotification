package model

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()" db:"id"`
	Timestamp   time.Time `json:"timestamp" gorm:"not null" db:"timestamp"`
	Platform    string    `json:"platform" gorm:"size:20;not null" db:"platform"`
	Description string    `json:"description" gorm:"size:255;not null" db:"description"`
}
