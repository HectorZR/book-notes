package main

import (
	"book-notes-app/book-notes/models"
	"book-notes-app/book-notes/seeders"
	"image/color"
	"log"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.RunDatabase()

	if os.Getenv("SEED") == "true" {
		seeders.RunSeeders()
		return
	}

	runApp()
}

func runApp() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	container := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(), text4, layout.NewSpacer())

	button := widget.NewButton("Click me!", func() {
		text1 = canvas.NewText("Not Hello", color.White)
		container = fyne.NewContainerWithLayout(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)
		myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), container, centered))
	})

	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), container, centered, button))

	myWindow.ShowAndRun()
}
