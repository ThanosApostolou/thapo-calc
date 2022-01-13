package main

import (
	"thapo-calc-cli/command"
	"thapo-calc-lib/libservice"
)

func main() {
	var services = libservice.InitializeServices()
	command.InitializeCommands(services)
}
