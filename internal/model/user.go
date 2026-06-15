package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string         `gorm:"type:varchar(64);uniqueIndex;not null" json:"username"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	Nickname     string         `gorm:"type:varchar(64)" json:"nickname"`
	Email        string         `gorm:"type:varchar(128)" json:"email"`
	Phone        string         `gorm:"type:varchar(32)" json:"phone"`
	Avatar       string         `gorm:"type:varchar(255)" json:"avatar"`
	Status       int            `gorm:"type:tinyint;default:1" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}
