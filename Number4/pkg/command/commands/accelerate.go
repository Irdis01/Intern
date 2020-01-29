package commands

type accelerateCommand struct {
	reciver car
}

func (a *accelerateCommand) Execute() (err error) {
	err = a.reciver.Accelerate()
	return
}

func (a *accelerateCommand) Cancel() (err error) {
	err = a.reciver.Slowdown()
	return
}

// NewAccelerateCommand конструктор команды для ускорения
func NewAccelerateCommand(reciever car) Commands {
	return &accelerateCommand{
		reciver: reciever,
	}
}
