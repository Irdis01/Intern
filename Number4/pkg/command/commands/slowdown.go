package commands

type slowDownCommand struct {
	carReciver car
}

func (a *slowDownCommand) Execute() (err error) {
	err = a.carReciver.Slowdown()
	return
}

func (a *slowDownCommand) Cancel() (err error) {
	err = a.carReciver.Accelerate()
	return
}

// NewSlowDownCommand конструктор команды для замедления
func NewSlowDownCommand(reciever car) Commands {
	return &slowDownCommand{
		carReciver: reciever,
	}
}
