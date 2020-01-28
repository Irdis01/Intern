package main

import (
	"github.com/Irdis01/Intern/Number1/pkg/chandelier"
	"github.com/Irdis01/Intern/Number1/pkg/crane"
	"github.com/Irdis01/Intern/Number1/pkg/facade"
	"github.com/Irdis01/Intern/Number1/pkg/microwave"
	"github.com/Irdis01/Intern/Number1/pkg/models"
)

func main() {
	ownerFood := models.Food{Temp: 10}
	ownerLamp := models.Lamp{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	}
	ownerBath := models.Bath{
		Volume:    0,
		MaxVolume: 10,
	}
	smartCrane := crane.NewCrane(&ownerBath)
	smartChandelier := chandelier.NewChandelier()
	smartChandelier.Add(&ownerLamp)
	smartMicrowave := microwave.NewMicrowave()
	smartMicrowave.Add(&ownerFood)
	smartHouse := facade.NewFacade(smartCrane, smartMicrowave, smartChandelier)
	smartHouse.Prepare(30, 70)
}
