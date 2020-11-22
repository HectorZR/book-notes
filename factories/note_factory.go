package factories

import (
	"book-notes-app/book-notes/models"
	"log"
	"strconv"

	"github.com/icrowley/fake"
)

// NoteFactory returns a new note instance with fake data
func NoteFactory() *models.Note {
	note := new(models.Note)

	chapter, err := strconv.Atoi(fake.Digits())
	if err != nil {
		log.Fatalf("Error converting string to number")
	}

	page, err := strconv.Atoi(fake.Digits())
	if err != nil {
		log.Fatalf("Error converting string to number")
	}

	note.Chapter = uint(chapter)
	note.Page = uint(page)
	note.Note = fake.Paragraph()

	return note
}
