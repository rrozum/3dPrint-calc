package desktop

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (ui *Ui) render(page fyne.CanvasObject) {
	ui.mainWindow.SetContent(page)
}

func (ui *Ui) makeHome() fyne.CanvasObject {
	projects := createProjects(ui)

	return container.NewMax(projects)
}

func (ui *Ui) makeProjectPage(name string) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Project Page for - "+name),
		container.NewAdaptiveGrid(3, container.NewWithoutLayout(), widget.NewButton("Back to projects", func() {
			ui.render(ui.makeHome())
		}), container.NewWithoutLayout()),
	)
}

func createProjects(ui *Ui) fyne.CanvasObject {
	projectsLayout := &ProjectLayout{}

	projectsContainer := container.New(projectsLayout)

	var rows []fyne.CanvasObject

	projectList, _ := ui.Services.Projects.GetAll()

	for _, project := range projectList {
		rows = append(rows, NewProject(project, projectsContainer, ui))
	}

	projectsContainer.Objects = rows

	scroll := container.NewVScroll(container.NewMax(projectsContainer))
	scroll.SetMinSize(fyne.NewSize(400, 200))

	titles := container.New(
		projectsLayout,
		container.NewVBox(
			widget.NewRichTextWithText("Список проектов"),
		),
	)

	return container.NewCenter(container.NewGridWrap(fyne.NewSize(600, 300), container.NewVBox(titles, scroll)))
}
