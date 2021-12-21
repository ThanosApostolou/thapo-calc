package command

import (
	"flag"
	"fmt"
)

type DefaultCommand struct {
	helpFlag *bool
}

var _ CommandI = (*DefaultCommand)(nil) // implement interface

func NewDefaultCommand() *DefaultCommand {
	var helpFlag = flag.Bool("help", false, "help message")
	return &DefaultCommand{
		helpFlag: helpFlag,
	}
}

func (this *DefaultCommand) Run() {
	flag.Parse()
	if len(flag.Args()) != 0 {
		panic("Wrong command")
	}

	if *this.helpFlag {
		this.Help()
	}
}

// Run Help func
func (this *DefaultCommand) Help() {
	fmt.Println("default command help message")
}
