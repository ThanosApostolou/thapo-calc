package main

import (
	"thapo-calc-cli/command"
	"thapo-calc-cli/service"
)

func main() {
	var services = service.InitializeServices()
	command.InitializeCommands(services)
}
