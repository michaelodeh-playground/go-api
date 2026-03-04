package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID    string `gorm:"primaryKey" json:"id"` // Primary key
	Name  string `json:"name"`
	Email string `gorm:"unique" json:"email"` // Unique column
	Age   int    `json:"age"`
}

func (u *Users) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}
