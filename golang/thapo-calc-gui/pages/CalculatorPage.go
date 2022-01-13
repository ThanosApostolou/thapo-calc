package pages

import (
	"fmt"
	"strconv"
	"thapo-calc-gui/service"
	"thapo-calc-lib/libservice"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type CalculatorPage struct {
	CalcService *libservice.CalcService
	GuiService  *service.GuiService

	State *CalculatorPageState
	El    *CalculatorPageEl
}

type CalculatorPageState struct {
}

type CalculatorPageEl struct {
	Container   *fyne.Container
	InputEntry  *widget.Entry
	ResultEntry *widget.Entry
}

func NewCalculatorPage(calcService *libservice.CalcService, guiService *service.GuiService) *CalculatorPage {
	var calculatorPage CalculatorPage
	var equalButton = widget.NewButton("=", calculatorPage.calcResult)

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
	page.Add(equalButton)

	var el = CalculatorPageEl{
		Container:   page,
		InputEntry:  inputEntry,
		ResultEntry: resultEntry,
	}

	calculatorPage = CalculatorPage{
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

func (cpage *CalculatorPage) calcResult() {
	fmt.Println("CalculatorPage calcResult")
	res, _ := cpage.CalcService.Calculate(cpage.El.InputEntry.Text)
	cpage.El.ResultEntry.SetText(strconv.FormatFloat(res, 'f', -1, 64))
	// cpage.El.ResultEntry.SetText(fmt.Sprintf("%f", res))
}
