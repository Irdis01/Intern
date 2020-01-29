package commands

type stopCommand struct {
	reciver car
}

func (s *stopCommand) Execute() (err error) {
	s.reciver.Slowdown()
	err = s.reciver.StopEngine()
	if err == nil {
		err = s.reciver.TurnOnHandbrake()
		return
	}
	s.reciver.Accelerate()
	return
}

func (s *stopCommand) Cancel() (err error) {
	s.reciver.StartEngine()
	s.reciver.TurnOffHandbrake()
	err = s.reciver.StartMove()
	return
}

// NewStopCommand конструктор команды для остановки машины
func NewStopCommand(reciver car) Commands {
	return &stopCommand{
		reciver: reciver,
	}
}
