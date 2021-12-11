package commands

import "os"

type CommandI interface {
	Run()
}

func InitializeCommands() {
	if len(os.Args) > 1 && os.Args[1] == CALC_COMMAND {
		var calcCommand = NewCalcCommand()
		calcCommand.Run()
	} else {
		var defaultCommand = NewDefaultCommand()
		defaultCommand.Run()
	}
}
