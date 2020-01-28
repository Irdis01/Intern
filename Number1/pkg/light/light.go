package light

type Light interface {
	TurnOn()
}

type light struct {
	state bool
}

func (l *light) TurnOn() {
	l.state = true
}

// NewLight return new light
func NewLight() Light {
	return &light{}
}
