package facade

type bath interface {
	FillUp(volume int)
}

type light interface {
	TurnOn()
}

type microwave interface {
	StartHeat(temp int)
}

// Manager интерфейс фасада
type Facade interface {
	PrepForOwner(waterTemp int, foodTemp int)
}

type facade struct {
	ownerBath      bath
	ownerLight     light
	ownerMicrowave microwave
}

func (f *facade) PrepForOwner(waterVolume int, foodTemp int) {
	f.ownerBath.FillUp(waterVolume)
	f.ownerLight.TurnOn()
	f.ownerMicrowave.StartHeat(foodTemp)
}

func NewFacade(ownerBath bath, ownerMicrowave microwave, ownerLight light) Facade {
	return &facade{
		ownerMicrowave: ownerMicrowave,
		ownerLight:     ownerLight,
		ownerBath:      ownerBath,
	}
}
