package entity

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username  string     `gorm:"size:255;not null;" json:"username"`
	Password  string     `gorm:"size:255;not null;" json:"-"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Reviews   []Review   `gorm:"foreignkey:AccountID" json:"reviews"`
	Favorites []Favorite `gorm:"foreignkey:AccountID" json:"favorites"`
}
