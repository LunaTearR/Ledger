package models

import (
	"time"

	"gorm.io/gorm"
)

type AdminLog struct {
	gorm.Model
	AdminID         User      `gorm:"foreignKey:UserID" json:"-"`
	Action          string    `gorm:"column:action;size:255;not null"`
	Target_user     User      `gorm:"foreignKey:UserID" json:"-"`
	ActionTimestamp time.Time `gorm:"column:action_timestamp;default:current_timestamp"`
}
