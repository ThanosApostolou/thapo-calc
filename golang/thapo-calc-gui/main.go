package main

import (
	"fmt"
	gservice "thapo-calc-gui/service"
	"thapo-calc-lib/service"
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
	CalcService *service.CalcService
	GuiService  *gservice.GuiService
}

func (ss *Services) GetCalcService() *service.CalcService {
	return ss.CalcService
}

func InitializeServices() *Services {
	var calcService = service.NewCalcService()
	var guiService = gservice.NewGuiService()
	// fmt.Println("res2", calcService.CalcExpression(""))
	return &Services{
		CalcService: calcService,
		GuiService:  guiService,
	}
}
