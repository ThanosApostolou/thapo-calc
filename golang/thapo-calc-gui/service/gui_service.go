package service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type GuiService struct {
	MyApp *fyne.App
}

func NewGuiService() *GuiService {
	return &GuiService{}
}

func (gs *GuiService) CreateApp() {
	var myapp = app.New()
	gs.MyApp = &myapp
}

func (gs *GuiService) RunApp() {
	(*gs.MyApp).Run()
}

func (gs *GuiService) CreateWindow(title string) fyne.Window {
	return (*gs.MyApp).NewWindow(title)
}
