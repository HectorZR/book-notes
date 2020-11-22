package factories

import (
	"book-notes-app/book-notes/models"
	"math/rand"
	"time"

	"github.com/icrowley/fake"
)

// BookFactory returns a new book instance with fake data
func BookFactory() *models.Book {
	book := new(models.Book)

	book.Name = fake.Title()
	book.Status = uint8(rand.Intn(models.READ))
	book.StartDate = time.Now()
	book.EndDate = time.Now()
	book.Notes = []models.Note{
		{Page: NoteFactory().Page, Chapter: NoteFactory().Chapter, Note: NoteFactory().Note},
		{Page: NoteFactory().Page, Chapter: NoteFactory().Chapter, Note: NoteFactory().Note},
		{Page: NoteFactory().Page, Chapter: NoteFactory().Chapter, Note: NoteFactory().Note},
	}

	return book
}
