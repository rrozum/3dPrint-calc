package desktop

import (
	"3dPrintCalc/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"time"
)

func Run(services *service.Services) {
	application := app.New()
	w := application.NewWindow("Hello World")
	w.Resize(fyne.Size{Width: 500, Height: 100})

	content := container.New(layout.NewVBoxLayout())

	updateData(content, services)

	w.SetContent(content)

	go func() {
		for range time.Tick(time.Second) {
			updateData(content, services)
		}
	}()

	w.ShowAndRun()
}

func updateData(content *fyne.Container, services *service.Services) {
	content.RemoveAll()

	applicationSettings, _ := services.ApplicationSettings.GetAll()
	for _, applicationSetting := range applicationSettings {
		settingString := applicationSetting.Key + ": " + applicationSetting.Value

		content.Add(canvas.NewText(settingString, color.White))
	}

	content.Refresh()
}
