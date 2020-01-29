package commands

type startCommand struct {
	reciever car
}

func (s *startCommand) Execute() (err error) {
	s.reciever.StartEngine()
	s.reciever.TurnOffHandbrake()
	err = s.reciever.StartMove()
	return
}

func (s *startCommand) Cancel() (err error) {
	s.reciever.Slowdown()
	err = s.reciever.StopEngine()
	if err != nil {
		err = s.reciever.TurnOnHandbrake()
	}
	return
}

// NewStartCommand конструктор команды для запуска машины
func NewStartCommand(reciver car) Commands {
	return &startCommand{
		reciever: reciver,
	}
}
