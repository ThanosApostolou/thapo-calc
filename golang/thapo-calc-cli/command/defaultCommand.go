package command

import (
	"flag"
	"fmt"
	"log"
	"thapo-calc-cli/service"
)

type DefaultCommand struct {
	calcService *service.CalcService
	helpFlag    *bool
}

var _ CommandI = (*DefaultCommand)(nil) // implement interface

func NewDefaultCommand(calcService *service.CalcService) *DefaultCommand {
	var helpFlag = flag.Bool("help", false, "help message")
	return &DefaultCommand{
		calcService: calcService,
		helpFlag:    helpFlag,
	}
}

func (this *DefaultCommand) Run() {
	flag.Parse()
	if len(flag.Args()) != 0 {
		panic("Wrong command")
	}

	if *this.helpFlag {
		this.Help()
		return
	}

	this.Tui()

}

// Run Help func
func (this *DefaultCommand) Help() {
	fmt.Println("default command help message")
}

func (this *DefaultCommand) Tui() {
	fmt.Println("Give an expression or \"exit\" to exit (supported operators +,-,*,/,^):")
	for true {
		var input string
		fmt.Print("> ")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		}
		result, err := this.calcService.Calculate(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("=", result)
		fmt.Println()
	}
}
