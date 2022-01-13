package command

import (
	"flag"
	"fmt"
	"log"
	"os"
	"thapo-calc-lib/libservice"
)

const CALC_COMMAND = "calc"

type CalcCommand struct {
	calcService *libservice.CalcService // injected
	Command     string
	cmdFlagSet  *flag.FlagSet
	helpFlag    *bool
}

var _ CommandI = (*CalcCommand)(nil) // implement interface

func NewCalcCommand(calcService *libservice.CalcService) *CalcCommand {
	var cmdFlagSet = flag.NewFlagSet(CALC_COMMAND, flag.ContinueOnError)
	var helpFlag = cmdFlagSet.Bool("help", false, "help message for calc")
	return &CalcCommand{
		calcService: calcService,
		Command:     CALC_COMMAND,
		cmdFlagSet:  cmdFlagSet,
		helpFlag:    helpFlag,
	}
}

func (this *CalcCommand) Run() {
	this.cmdFlagSet.Parse(os.Args[2:])
	if len(os.Args) > 0 && os.Args[1] == this.Command {
		if *this.helpFlag {
			this.Help()
			return
		}

		this.Calc()
	}
}

func (this *CalcCommand) Help() {
	fmt.Println("calc command help message")
}

func (this *CalcCommand) Calc() {
	if len(this.cmdFlagSet.Args()) != 1 {
		this.Help()
		os.Exit(1)
	}
	var expression = this.cmdFlagSet.Args()[0]
	var result, err = this.calcService.Calculate(expression)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result)
}
