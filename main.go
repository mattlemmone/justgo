package main

import (

	// "github.com/leaanthony/mewn"
	// "github.com/wailsapp/wails"
  // "fmt"
  // "os/exec"
	"os"
  "github.com/therecipe/qt/widgets"
  "github.com/mattlemmone/popo/search"
  "github.com/mattlemmone/popo/desktop"
)

var apps []*desktop.Application
const MaxResults = 10

func main() {
  // environment := desktop.Linux{}

  // println("Scanning...")
  // apps = environment.DesktopApplications()
  // userFiles := environment.UserFiles()

  // results := search.FuzzyFindFile(text[:len(text)-1], userFiles)

  // topResult := results[0]
  // environment.LaunchApplication(topResult)

	// js := mewn.String("./frontend/dist/app.js")
	// css := mewn.String("./frontend/dist/index.css")

	// app := wails.CreateApp(&wails.AppConfig{
	// 	Width:  1024,
	// 	Height: 200,
	// 	Title:  "popo",
	// 	JS:     js,
	// 	CSS:    css,
	// })

	// app.Bind(onTextChange)
	// app.Bind(&environment)
	// app.Run()

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create a window
	// with a minimum size of 250*200
	// and sets the title to "Hello Widgets Example"
	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Hello Widgets Example")

	// create a regular widget
	// give it a QVBoxLayout
	// and make it the central widget of the window
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	// create a line edit
	// with a custom placeholder text
	// and add it to the central widgets layout
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Write something ...")
	widget.Layout().AddWidget(input)

	// create a button
	// connect the clicked signal
	// and add it to the central widgets layout
	button := widgets.NewQPushButton2("and click me!", nil)
	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	widget.Layout().AddWidget(button)

	// make the window visible
	window.Show()

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}

func onTextChange(newText string) []*desktop.Application{
  results := search.FuzzyFindApplication(newText, apps)
  numResults := min(len(results), MaxResults)

  return results[:numResults]
}


func min(a int, b int) int {
  if a > b {
    return b
  }

  return a
}
