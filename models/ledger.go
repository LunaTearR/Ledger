package models

import (
	"time"

	"gorm.io/gorm"
)

type Ledger struct {
	gorm.Model
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
	Title       string    `gorm:"type:varchar(200);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	Amount      float64   `gorm:"type:numeric(15,2);not null" json:"amount"`
	Type        string    `gorm:"type:varchar(20);not null" json:"type"`
	Category    string    `gorm:"type:varchar(100)" json:"category,omitempty"`
	Date        time.Time `gorm:"not null" json:"date"`
}