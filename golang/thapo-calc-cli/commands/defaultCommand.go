package commands

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

	fmt.Println("default command")
	fmt.Println("default command", *this.helpFlag)
	if *this.helpFlag {
		this.RunHelp()
	}
}

// Run Help func
func (this *DefaultCommand) RunHelp() {
	fmt.Println("default command help message")
}
