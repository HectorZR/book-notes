package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// RunDatabase gets a connection and run all models
func RunDatabase() {
	db := GetConnection()

	db.AutoMigrate(&Book{}, &Note{})

	// some testing data
	// db.Create(&Book{Name: "Hello book", Status: 0, StartDate: time.Now()})
}

// GetConnection open a new connection with database
func GetConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("book_notes.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
