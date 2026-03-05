package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        string         `gorm:"primaryKey" json:"id"` // Primary key
	Name      string         `json:"name"`
	Email     string         `gorm:"unique" json:"email"` // Unique column
	Age       int            `gorm:"default:18" json:"age"`
	Balance   float64        `gorm:"default:0.0" json:"balance"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}
