package commands

type stopCommand struct {
	carReciever car
}

func (s *stopCommand) Execute() (err error) {
	s.carReciever.Slowdown()
	err = s.carReciever.StopEngine()
	if err == nil {
		err = s.carReciever.TurnOnHandbrake()
		return
	}
	s.carReciever.Accelerate()
	return
}

func (s *stopCommand) Cancel() (err error) {
	s.carReciever.StartEngine()
	s.carReciever.TurnOffHandbrake()
	err = s.carReciever.StartMove()
	return
}

// NewStopCommand конструктор команды для остановки машины
func NewStopCommand(reciver car) Commands {
	return &stopCommand{
		carReciever: reciver,
	}
}
