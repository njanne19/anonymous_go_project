package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

// should only call call app .run and ONCE
// can maybe create multiple windows?
func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	label := widget.NewLabel("Hello World 1")

	w.Resize(fyne.NewSize(200, 200))

	updateTime(label)

	w.SetContent(container.NewVBox(
		label,
		widget.NewEntry(),
	))

	w.SetMaster()
	w.Show()

	w2 := a.NewWindow("Larger")

	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))

	w2.Resize(fyne.NewSize(400, 400))
	w2.Show()

	go func() {
		for range time.Tick(time.Second) {
			updateTime(label)
		}
	}()

	a.Run()
}
