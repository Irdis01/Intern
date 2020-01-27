package commands

type startCommand struct {
	carReciver car
}

func (s *startCommand) Execute() (err error) {
	s.carReciver.StartEngine()
	s.carReciver.TurnOffHandbrake()
	err = s.carReciver.StartMove()
	return
}

func (s *startCommand) Cancel() (err error) {
	s.carReciver.Slowdown()
	err = s.carReciver.StopEngine()
	if err != nil {
		err = s.carReciver.TurnOnHandbrake()
	}
	return
}

// NewStartCommand конструктор команды для запуска машины
func NewStartCommand(reciver car) Commands {
	return &startCommand{
		carReciver: reciver,
	}
}
