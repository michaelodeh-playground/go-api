package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transactions struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	Amount      float64        `json:"amount"`
	UserID      string         `json:"user_id"`
	User        Users          `gorm:"foreignKey:UserID" json:"user"`
	ReferenceID string         `gorm:"uniqueIndex;not null" json:"reference_id"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (t *Transactions) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return nil
}
