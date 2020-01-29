package main

import (
	"log"

	"github.com/Irdis01/Intern/Number4/pkg/command/car"
	"github.com/Irdis01/Intern/Number4/pkg/command/commands"
	"github.com/Irdis01/Intern/Number4/pkg/command/palyer"
)

func main() {
	var (
		commandsError error
	)
	carDriver := palyer.NewDriver()
	hisCar := car.NewCar()
	carDriver.StoreCommand(commands.NewStartCommand(hisCar))
	carDriver.StoreCommand(commands.NewAccelerateCommand(hisCar))
	carDriver.StoreCommand(commands.NewStopCommand(hisCar))
	commandsError = carDriver.Execute()
	log.Println(hisCar.GetSpeed())
	log.Println(commandsError)
	commandsError = carDriver.Execute()
	log.Println(hisCar.GetSpeed())
	log.Println(commandsError)
	commandsError = carDriver.Cancel()
	log.Println(hisCar.GetSpeed())
	log.Println(commandsError)
	commandsError = carDriver.Execute()
	log.Println(hisCar.GetSpeed())
	log.Println(commandsError)
	commandsError = carDriver.Execute()
	log.Println(hisCar.GetSpeed())
	log.Println(commandsError)

}
