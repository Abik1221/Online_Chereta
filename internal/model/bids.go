package models

import "gorm.io/gorm"

type Bid struct {
    gorm.Model
    UserID    uint    `gorm:"not null"`
    ItemID    uint    `gorm:"not null"`
    BidAmount float64 `gorm:"not null"`
    Status    string  `gorm:"default:'active'"` // e.g., "active", "won", "lost"
}