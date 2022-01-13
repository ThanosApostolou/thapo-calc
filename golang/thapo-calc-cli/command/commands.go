package command

import (
	"os"
	"thapo-calc-lib/libservice"
)

type CommandI interface {
	Run()
}

func InitializeCommands(services *libservice.Services) {
	if len(os.Args) > 1 && os.Args[1] == CALC_COMMAND {
		var calcCommand = NewCalcCommand(services.GetCalcService())
		calcCommand.Run()
	} else {
		var defaultCommand = NewDefaultCommand(services.GetCalcService())
		defaultCommand.Run()
	}
}
