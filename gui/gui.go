package gui

import (
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

func Build() {
	app := app.New()

	w := app.NewWindow("Hello")

	box := widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewEntry(),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	)

	w.SetContent(box)
	w.ShowAndRun()
}
