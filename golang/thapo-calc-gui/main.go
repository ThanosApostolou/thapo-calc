package main

import (
	"fmt"
	"thapo-calc-lib/service"
)

func main() {
	var services = service.InitializeServices()
	fmt.Println(services.GetCalcService().Calculate("2+3/2"))
	var appComp = NewAppComp()
	appComp.Init()
}
