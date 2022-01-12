package components

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type HeaderComp struct {
	El *widget.Toolbar
}

func NewHeaderComp() *HeaderComp {
	var headerComp = HeaderComp{}
	// headerComp.El = container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewLabel("test"))
	var menuAction = widget.NewToolbarAction(theme.MenuIcon(), func() {

	})
	headerComp.El = widget.NewToolbar(widget.NewToolbarSpacer(), menuAction)
	return &headerComp
}

func (thisHC *HeaderComp) OnInit() {}
