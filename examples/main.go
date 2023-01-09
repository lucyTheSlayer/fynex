package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/lucyTheSlayer/fynex"
)

func main() {
	a := app.New()
	w := a.NewWindow("PaddingLayout")
	text1 := widget.NewLabel("topleft")
	w.SetContent(container.New(fynex.NewPaddingLayout(100, 10, 10, 100), text1))
	w.ShowAndRun()
}
