package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
