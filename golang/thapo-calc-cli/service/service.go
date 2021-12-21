package service

type Services struct {
	calcService *CalcService
}

func (this *Services) GetCalcService() *CalcService {
	return this.calcService
}

func InitializeServices() *Services {
	var calcService = NewCalcService()
	// fmt.Println("res2", calcService.CalcExpression(""))
	return &Services{
		calcService: calcService,
	}
}
