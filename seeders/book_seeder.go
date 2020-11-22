package seeders

import (
	"book-notes-app/book-notes/factories"
)

// RunBookSeeder runs book seeder
func RunBookSeeder(times int) {
	for i := 0; i < times; i++ {
		seed := factories.BookFactory()
		result := seed.Create()

		if result.Error != nil {
			panic("Failed to run book seeder")
		}
	}
}
