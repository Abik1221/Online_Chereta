package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username     string `gorm:"unique;not null"`
    Email        string `gorm:"unique;not null"`
    PasswordHash string `gorm:"not null"`
    GoogleID     string `gorm:"unique"`
    PhoneNumber  string
    Role         string `gorm:"default:user"`
}