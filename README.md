## PaddingLayout
```go
func main() {
  a := app.New()
  w := a.NewWindow("PaddingLayout")
  text1 := widget.NewLabel("topleft")
  w.SetContent(container.New(fynex.NewPaddingLayout(100, 10, 10, 100), text1))
  w.ShowAndRun()
}
```