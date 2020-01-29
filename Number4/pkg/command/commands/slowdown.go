package commands

type slowDownCommand struct {
	reciver car
}

func (a *slowDownCommand) Execute() (err error) {
	err = a.reciver.Slowdown()
	return
}

func (a *slowDownCommand) Cancel() (err error) {
	err = a.reciver.Accelerate()
	return
}

// NewSlowDownCommand конструктор команды для замедления
func NewSlowDownCommand(reciver car) Commands {
	return &slowDownCommand{
		reciver: reciver,
	}
}
