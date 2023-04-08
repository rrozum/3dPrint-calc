package desktop

import (
	"3dPrintCalc/internal/domain"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type project struct {
	widget.BaseWidget
	project           domain.Project
	projectsContainer *fyne.Container

	button *widget.Button
	dots   *fyne.Container

	openProjectFunction func()
}

func NewProject(p domain.Project, projectsContainer *fyne.Container, ui *Ui) *project {
	project := &project{project: p, projectsContainer: projectsContainer}

	project.ExtendBaseWidget(project)

	project.openProjectFunction = func() {
		ui.render(ui.makeProjectPage(p))
	}

	menu := fyne.NewMenu("",
		fyne.NewMenuItem("Открыть", project.openProjectFunction),
		fyne.NewMenuItem("Удалить проект", func() {
			project.remove(projectsContainer)
		}),
	)

	project.button = widget.NewButtonWithIcon("", theme.MoreHorizontalIcon(), func() {
		position := fyne.CurrentApp().Driver().AbsolutePositionForObject(project.button)
		position.Y += project.button.Size().Height

		widget.ShowPopUpMenuAtPosition(menu, fyne.CurrentApp().Driver().CanvasForObject(project.button), position)
	})

	project.button.Importance = widget.LowImportance

	project.dots = container.NewHBox(widget.NewSeparator(), container.NewVBox(project.button))

	return project
}

func (p *project) remove(projectsContainer *fyne.Container) {
	p.removeProjectFromContainer(projectsContainer)
}

func (p *project) removeProjectFromContainer(projectsContainer *fyne.Container) {
	for j := 0; j < len(projectsContainer.Objects); j++ {
		if p.project.Id == projectsContainer.Objects[j].(*project).project.Id {
			projectsContainer.Remove(projectsContainer.Objects[j])
			break
		}
	}
}

func (p *project) CreateRenderer() fyne.WidgetRenderer {
	widgetBtn := widget.NewButton("", p.openProjectFunction)

	btn := container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.NRGBA{A: 0x59}), widgetBtn)
	projectName := widget.NewRichTextWithText(p.project.Name)
	projectPrice := widget.NewRichTextWithText(fmt.Sprint(p.project.Price) + " ₽")

	c := container.NewMax(btn,
		container.NewBorder(nil, nil, container.NewHBox(projectName), container.NewHBox(projectPrice, p.dots)))

	return widget.NewSimpleRenderer(c)
}
