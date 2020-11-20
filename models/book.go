package models

import (
	"time"

	"gorm.io/gorm"
)

// Book structure
type Book struct {
	gorm.Model
	Name      string    `gorm:"unique;size:100;not null"`
	Status    uint8     `gorm:"not null"`
	StartDate time.Time `gorm:"autoCreateTime;not null"`
	EndDate   time.Time
	Notes     []Note
}

// Create a new book record
// func (b *Book) Create(book Book) bool {
// 	db := GetConnection()

// 	db.Create(&Book{book})

// 	return true
// }
