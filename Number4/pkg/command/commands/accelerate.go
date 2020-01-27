package commands

type accelerateCommand struct {
	carReciver car
}

func (a *accelerateCommand) Execute() (err error) {
	err = a.carReciver.Accelerate()
	return
}

func (a *accelerateCommand) Cancel() (err error) {
	err = a.carReciver.Slowdown()
	return
}

// NewAccelerateCommand конструктор команды для ускорения
func NewAccelerateCommand(reciever car) Commands {
	return &accelerateCommand{
		carReciver: reciever,
	}
}
