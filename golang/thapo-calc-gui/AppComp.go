package main

import (
	"fmt"
	"thapo-calc-gui/components"
	"thapo-calc-gui/pages"
	"thapo-calc-lib/libservice"

	"thapo-calc-gui/service"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type AppComp struct {
	CalcService *libservice.CalcService
	GuiService  *service.GuiService

	State *AppCompState
	El    *AppCompEl
}

type AppCompState struct {
}

type AppCompEl struct {
	Window     fyne.Window
	WindowGrid *fyne.Container
	HeaderComp *components.HeaderComp
	Page       *pages.CalculatorPage
}

func NewAppComp(calcService *libservice.CalcService, guiService *service.GuiService) *AppComp {
	var page = pages.NewCalculatorPage(calcService, guiService)
	var headerComp = components.NewHeaderComp()
	var windowGrid = container.NewBorder(headerComp.El, nil, nil, nil, page.El.Container)

	var window = guiService.CreateWindow("ThApo Calculator")
	window.SetContent(windowGrid)

	var el = AppCompEl{
		Window:     window,
		WindowGrid: windowGrid,
		HeaderComp: headerComp,
		Page:       page,
	}

	var appComp = &AppComp{
		CalcService: calcService,
		GuiService:  guiService,

		State: &AppCompState{},
		El:    &el,
	}
	appComp.El.Window.Resize(fyne.NewSize(400, 400))

	return appComp
}

func (appcomp *AppComp) OnStarted() {
	fmt.Println("AppComp OnStarted")
	appcomp.El.Window.Show()
	appcomp.El.Page.OnStarted()
}

func (appcomp *AppComp) OnStopped() {
	fmt.Println("AppComp OnStopped")
	appcomp.El.Page.OnStopped()
}
