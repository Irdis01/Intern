package main

import (
	"github.com/Irdis01/Intern/Number1/pkg/bath"
	"github.com/Irdis01/Intern/Number1/pkg/facade"
	"github.com/Irdis01/Intern/Number1/pkg/light"
	"github.com/Irdis01/Intern/Number1/pkg/microwave"
	"github.com/Irdis01/Intern/Number1/pkg/models"
)

func main() {
	ownerFood := models.Food{Temp: 10}
	smartBath := bath.NewBath()
	smartLamp := light.NewLight()
	smartMicrowave := microwave.NewMicrowave(ownerFood)
	smartHouse := facade.NewFacade(smartBath, smartMicrowave, smartLamp)
	smartHouse.PrepForOwner(30, 70)
}
