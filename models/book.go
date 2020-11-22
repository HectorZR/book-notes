package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	NOT_READING = 0
	READING     = 1
	READ        = 2
)

// Book structure
type Book struct {
	gorm.Model
	Name      string    `gorm:"unique;size:100;not null"`
	Status    uint8     `gorm:"not null;default:0"`
	StartDate time.Time `gorm:"autoCreateTime;not null"`
	EndDate   time.Time
	Notes     []Note
}

// Create a new book record
func (b *Book) Create() *gorm.DB {
	db := GetConnection()

	return db.Create(&b)
}
