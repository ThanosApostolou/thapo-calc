package commands

import (
	"flag"
	"fmt"
	"os"
)

const CALC_COMMAND = "calc"

type CalcCommand struct {
	Command    string
	cmdFlagSet *flag.FlagSet
	helpFlag   *bool
}

var _ CommandI = (*CalcCommand)(nil) // implement interface

func NewCalcCommand() *CalcCommand {
	var cmdFlagSet = flag.NewFlagSet(CALC_COMMAND, flag.PanicOnError)
	var helpFlag = cmdFlagSet.Bool("help", false, "help message for calc")
	return &CalcCommand{
		Command:    CALC_COMMAND,
		cmdFlagSet: cmdFlagSet,
		helpFlag:   helpFlag,
	}
}

func (this *CalcCommand) Run() {
	this.cmdFlagSet.Parse(os.Args[2:])
	if len(os.Args) > 0 && os.Args[1] == this.Command {
		fmt.Println("calc command", this.cmdFlagSet.Args())
		fmt.Println("calc command", *this.helpFlag)
		if *this.helpFlag {
			fmt.Println("calc command help message")
		}
	}
}
