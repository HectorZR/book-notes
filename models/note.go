package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Page    uint
	Chapter uint
	Note    string
	BookID  uint
	Book    Book `gorm:"constraint:OnDelete:CASCADE"`
}

// Create makes a new note in database
func (n *Note) Create() *gorm.DB {
	db := GetConnection()

	return db.Create(&n)
}
