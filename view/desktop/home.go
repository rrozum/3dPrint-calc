package desktop

import (
	"3dPrintCalc/internal/domain"
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type detailParams struct {
	costTimeMachine  binding.String
	costTimeOperator binding.String
	costPlastic      binding.String
	costGeneral      binding.String
	costForClient    binding.String
}

func (ui *Ui) render(page fyne.CanvasObject) {
	ui.mainWindow.SetContent(page)
}

func (ui *Ui) makeHome() fyne.CanvasObject {
	projects := createProjects(ui)

	return container.NewMax(projects)
}

func (ui *Ui) makeProjectPage(project domain.Project) fyne.CanvasObject {
	header := canvas.NewText(project.Name, color.White)

	header.TextSize = 24
	header.TextStyle.Monospace = true

	headerProject := canvas.NewText("Параметры проекта", color.White)
	headerProject.TextSize = 20
	headerProject.TextStyle.Monospace = true

	headerDetail := canvas.NewText("Расчеты по детали", color.White)
	headerDetail.TextSize = 20
	headerDetail.TextStyle.Monospace = true

	details, _ := ui.Services.Details.GetByProject(project.Id)

	params := detailParams{
		costTimeMachine:  binding.NewString(),
		costTimeOperator: binding.NewString(),
		costPlastic:      binding.NewString(),
		costGeneral:      binding.NewString(),
		costForClient:    binding.NewString(),
	}

	params.costTimeMachine.Set("1234")
	params.costTimeOperator.Set("1234")
	params.costPlastic.Set("1234")
	params.costGeneral.Set("1234")
	params.costForClient.Set("1234")

	vBox := container.NewVBox(
		container.NewPadded(container.NewVBox(header)),
		container.NewMax(canvas.NewLine(color.White)),
		container.NewPadded(container.NewCenter(headerProject)),
		container.NewPadded(makeProjectParams(project).Objects...),
		container.NewMax(canvas.NewLine(color.White)),
		container.NewPadded(container.NewCenter(headerDetail)),
		container.NewPadded(makeDetailParams(params).Objects...),
	)

	setDetail := func(detail domain.Detail) {
		params.costTimeMachine.Set(detail.Name)
		params.costTimeOperator.Set(detail.Name)
		params.costPlastic.Set(detail.Name)
		params.costGeneral.Set(detail.Name)
		params.costForClient.Set(detail.Name)
	}

	backButton := widget.NewButton("Вернуться", func() {
		ui.render(ui.makeHome())
	})

	split := container.NewHSplit(makeNavForDetails(setDetail, details, backButton), vBox)
	split.Offset = 0.2

	return split
}

func makeDetailParams(params detailParams) fyne.Container {
	return *container.NewWithoutLayout(container.NewGridWithRows(
		2,
		container.NewGridWithColumns(2, widget.NewLabel("Параметр"), widget.NewLabel("Значение")),
		container.NewGridWithColumns(2, widget.NewLabel("Стоимость времени работы оборудования"), widget.NewLabelWithData(params.costTimeMachine)),
	))
}

func makeProjectParams(project domain.Project) fyne.Container {
	return *container.NewWithoutLayout(container.NewGridWithRows(
		3,
		container.NewGridWithColumns(2, widget.NewLabel("Параметр"), widget.NewLabel("Значение")),
		makeProjectParamEntry("Расход пластика (г) - Данные из ПО", "234"),
		makeProjectParamEntry("Время печати (часы) - Данные из ПО", "34"),
	))
}

func makeProjectParamEntry(paramName, entryDefaultValue string) *fyne.Container {
	entry := widget.NewEntry()
	entry.Text = entryDefaultValue
	return container.NewGridWithColumns(2, widget.NewLabel(paramName), entry)
}

func makeNavForDetails(setDetail func(detail domain.Detail), details []domain.Detail, backButton *widget.Button) fyne.CanvasObject {
	var detailIDs []string

	detailsMap := make(map[string]domain.Detail)

	for _, detail := range details {
		detailIDs = append(detailIDs, strconv.Itoa(detail.Id))
		detailsMap[strconv.Itoa(detail.Id)] = detail
	}

	tree := widget.NewTree(func(id widget.TreeNodeID) []widget.TreeNodeID {
		return detailIDs
	}, func(id widget.TreeNodeID) bool {
		return true
	}, func(b bool) fyne.CanvasObject {
		fmt.Println("Create node", b)
		return widget.NewLabel("Collection Widgets")
	}, func(id widget.TreeNodeID, b bool, object fyne.CanvasObject) {
		object.(*widget.Label).SetText(detailsMap[id].Name)
	})

	tree.OnSelected = func(uid widget.TreeNodeID) {
		setDetail(detailsMap[uid])
	}

	//a := fyne.CurrentApp()
	//themes := container.NewGridWithColumns(2,
	//	widget.NewButton("Dark", func() {
	//		a.Settings().SetTheme(theme.DarkTheme())
	//	}),
	//	widget.NewButton("Light", func() {
	//		a.Settings().SetTheme(theme.LightTheme())
	//	}),
	//)

	return container.NewBorder(widget.NewLabel("Делали проекта"), container.NewVBox(backButton), nil, nil, tree)
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
