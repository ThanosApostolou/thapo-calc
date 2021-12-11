package main

import "thapo-calc-cli/commands"

func main() {
	// cmd1FlagSet := flag.NewFlagSet("cmd1", flag.PanicOnError)
	// var _ = cmd1FlagSet.String("n", "", "help message for flag n")
	commands.InitializeCommands()
	// flag.Parse()
	// fmt.Println("Strarting cli: ", os.Args)
	// fmt.Println("flag.Args: ", flag.Args())
	// fmt.Println("nFlag: ", *nFlag)
	// var cmdFlagSet = flag.NewFlagSet("calc", flag.PanicOnError)
	// var helpFlag = cmdFlagSet.Bool("help", false, "help message for flag n")
	// cmdFlagSet.Parse(os.Args[2:])
	// fmt.Println("helpFlag", *helpFlag)
}
