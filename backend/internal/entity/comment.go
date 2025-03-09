package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `gorm:"not null;" json:"content"`
	AccountID uint   `gorm:"not null;" json:"account_id"` // Foreign key linking to the Account
	ReviewID  uint   `gorm:"not null;" json:"review_id"`  // Foreign key linking to the Review
}
