package pages

import (
	"fmt"
	gservice "thapo-calc-gui/service"
	"thapo-calc-lib/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type CalculatorPage struct {
	CalcService *service.CalcService
	GuiService  *gservice.GuiService

	State *CalculatorPageState
	El    *CalculatorPageEl
}

type CalculatorPageState struct {
}

type CalculatorPageEl struct {
	Container *fyne.Container
}

func NewCalculatorPage(calcService *service.CalcService, guiService *gservice.GuiService) *CalculatorPage {
	var helloLabel = widget.NewLabel("Hello Fyne!")
	var helloButton = widget.NewButton("Hi!", func() {
		helloLabel.SetText("Welcome :)")
	})

	var inputEntry = widget.NewEntry()
	inputEntry.SetPlaceHolder("Input")

	var resultEntry = widget.NewEntry()
	resultEntry.SetPlaceHolder("Result")
	resultEntry.Disable()
	var resultForm = container.New(layout.NewFormLayout())
	resultForm.Add(widget.NewLabel("="))
	resultForm.Add(resultEntry)

	var inputResultRow = container.NewGridWithColumns(2)
	inputResultRow.Add(inputEntry)
	inputResultRow.Add(resultForm)

	var page = container.NewVBox()
	page.Add(inputResultRow)
	page.Add(helloButton)

	var el = CalculatorPageEl{
		Container: page,
	}

	var calculatorPage = CalculatorPage{
		CalcService: calcService,
		GuiService:  guiService,

		State: &CalculatorPageState{},
		El:    &el,
	}
	return &calculatorPage
}

func (cpage *CalculatorPage) OnStarted() {
	fmt.Println("CalculatorPage OnStarted")
}

func (cpage *CalculatorPage) OnStopped() {
	fmt.Println("CalculatorPage OnStopped")
}
