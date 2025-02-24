package models

import "gorm.io/gorm"

type Item struct {
    gorm.Model
    Title        string  `gorm:"not null"`
    Description  string  `gorm:"not null"`
    StartingPrice float64 `gorm:"not null"`
    ReservePrice  float64 `gorm:"not null"`
    EndDate      int64   `gorm:"not null"` // Unix timestamp
    Condition    string  `gorm:"not null"`
    ImageURL     string  `gorm:"not null"`
    UserID       uint    `gorm:"not null"` // User who created the item
}


