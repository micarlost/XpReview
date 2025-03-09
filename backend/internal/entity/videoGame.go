package entity

import "gorm.io/gorm"

type VideoGame struct {
	gorm.Model
	Title       string     `gorm:"size:255;not null;" json:"title"`
	Description string     `gorm:"size:500;" json:"description"`
	ReleaseDate string     `gorm:"size:100;" json:"release_date"`
	Reviews     []Review   `gorm:"foreignkey:VideoGameID" json:"reviews"`   // One-to-many relationship with reviews
	Favorites   []Favorite `gorm:"foreignkey:VideoGameID" json:"favorites"` // Many-to-many relationship with favorites
}
