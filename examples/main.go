package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/lucyTheSlayer/fynex"
)

func main() {
	a := app.New()
	w := a.NewWindow("PaddingLayout")
	s := binding.NewString()
	s.Set("12.32")
	e := fynex.NewNumericalEntryWithData(s)
	w.SetContent(container.New(fynex.NewPaddingLayout(10, 10, 10, 10), e))
	w.ShowAndRun()
}
