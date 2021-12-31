package command

import (
	"os"
	"thapo-calc-lib/service"
)

type CommandI interface {
	Run()
}

func InitializeCommands(services *service.Services) {
	if len(os.Args) > 1 && os.Args[1] == CALC_COMMAND {
		var calcCommand = NewCalcCommand(services.GetCalcService())
		calcCommand.Run()
	} else {
		var defaultCommand = NewDefaultCommand(services.GetCalcService())
		defaultCommand.Run()
	}
}
