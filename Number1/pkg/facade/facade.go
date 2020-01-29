package facade

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
)

type crane interface {
	FillUp(volume int) (waterVolume int, err error)
}

type chandelier interface {
	TurnOn() (lampState bool, err error)
	TurnOff() (lampState bool, err error)
	Add(lamp *models.Lamp) (lampIsBurn bool, err error)
	Remove() (lamp *models.Lamp, err error)
}

type microwave interface {
	Heat(temp int) (foodTemp int, err error)
	Add(food *models.Food) (err error)
	Get() (foodInMicrowave *models.Food, err error)
}

// Facade facade interface
type Facade interface {
	Prepare(waterTemp int, foodTemp int) (bathResult int, lightResult bool, heatResult int, err error)
}

type facade struct {
	ownerBath      crane
	ownerLight     chandelier
	ownerMicrowave microwave
}

func (f *facade) Prepare(waterVolume int, foodTemp int) (bathResult int, lightResult bool, heatResult int, err error) {
	var (
		errBath,errLight, errHeat error
	)
	bathResult, errBath = f.ownerBath.FillUp(waterVolume)
	if errBath != nil {
		err=errBath
	}
	lightResult, errLight = f.ownerLight.TurnOn()
	if errLight != nil {
		err=errors.New(err.Error()+" "+errLight.Error() )
	}
	heatResult, errHeat = f.ownerMicrowave.Heat(foodTemp)
	if errHeat != nil {
		err=errors.New(err.Error()+" "+errHeat.Error() )
	}
	return
}

// NewFacade return new facade
func NewFacade(ownerBath crane, ownerMicrowave microwave, ownerLight chandelier) Facade {
	return &facade{
		ownerMicrowave: ownerMicrowave,
		ownerLight:     ownerLight,
		ownerBath:      ownerBath,
	}
}
