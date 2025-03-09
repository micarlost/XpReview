package entity

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	AccountID   uint `gorm:"not null;" json:"account_id"`    // Foreign key linking to Account
	VideoGameID uint `gorm:"not null;" json:"video_game_id"` // Foreign key linking to VideoGame
}
