package main

import (
	"fmt"
	"thapo-calc-gui/service"
	"thapo-calc-lib/libservice"
)

func main() {
	var services = InitializeServices()
	fmt.Println(services.GetCalcService().Calculate("2+3/2"))
	services.GuiService.CreateApp()
	var appComp = NewAppComp(services.CalcService, services.GuiService)
	(*services.GuiService.MyApp).Lifecycle().SetOnStarted(func() {
		appComp.OnStarted()
	})
	(*services.GuiService.MyApp).Lifecycle().SetOnStopped(func() {
		appComp.OnStopped()
	})
	services.GuiService.RunApp()
}

type Services struct {
	CalcService *libservice.CalcService
	GuiService  *service.GuiService
}

func (ss *Services) GetCalcService() *libservice.CalcService {
	return ss.CalcService
}

func InitializeServices() *Services {
	var calcService = libservice.NewCalcService()
	var guiService = service.NewGuiService()
	return &Services{
		CalcService: calcService,
		GuiService:  guiService,
	}
}
