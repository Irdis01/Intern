package facade

import "github.com/Irdis01/Intern/Number1/pkg/models"

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
	Prepare(waterTemp int, foodTemp int) (err error)
}

type facade struct {
	ownerBath      crane
	ownerLight     chandelier
	ownerMicrowave microwave
}

func (f *facade) Prepare(waterVolume int, foodTemp int) (err error) {
	_, err = f.ownerBath.FillUp(waterVolume)
	if err != nil {
		return
	}
	_, err = f.ownerLight.TurnOn()
	if err != nil {
		return
	}
	_, err = f.ownerMicrowave.Heat(foodTemp)
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
