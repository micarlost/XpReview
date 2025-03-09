package entity

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Content     string    `gorm:"not null;" json:"content"`
	Rating      float32   `gorm:"not null;" json:"rating"`        // For rating the game (e.g., 1-10)
	AccountID   uint      `gorm:"not null;" json:"account_id"`    // Foreign key linking to the Account
	VideoGameID uint      `gorm:"not null;" json:"video_game_id"` // Foreign key linking to the VideoGame
	Comments    []Comment `gorm:"foreignkey:ReviewID" json:"comments"`
}
