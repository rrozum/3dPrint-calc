package desktop

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Run(data string) {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel(data))
	w.ShowAndRun()
}
