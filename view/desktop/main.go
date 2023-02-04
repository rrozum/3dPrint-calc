package desktop

import (
	"3dPrintCalc/internal/service"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Ui struct {
	mainWindow fyne.Window
	Services   *service.Services
}

func Run(services *service.Services) {
	application := app.NewWithID("com.3dthing.printcalc")
	w := application.NewWindow("Калькулятор 3D печати")

	ui := &Ui{mainWindow: w, Services: services}

	w.Resize(fyne.NewSize(900, 600))
	ui.render(ui.makeHome())
	w.SetIcon(resourceIconPng)

	w.ShowAndRun()
}
