package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password     string    `gorm:"type:varchar(255);not null" json:"-"`
	Role         string    `gorm:"type:varchar(20);default:'user';not null" json:"role"`
	Status       bool      `gorm:"default:true" json:"status"`
	LastLogin    *time.Time `gorm:"column:last_login" json:"last_login"`
	TerminatedAt *time.Time `gorm:"column:terminated_at" json:"terminated_at"`
}