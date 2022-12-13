package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    int            `gorm:"primaryKey" json:"-"`
	Name      string         `gorm:"primaryKey" json:"name"`
	Email     string         `gorm:"primaryKey" json:"email"`
	Password  string         `gorm:"primaryKey" json:"password"`
}
