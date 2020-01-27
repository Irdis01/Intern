package car

import "errors"

// Car интерфейс автомобиля
type Car interface {
	StartEngine()
	StopEngine() error
	TurnOnHandbrake() error
	TurnOffHandbrake()
	StartMove() error
	Accelerate() error
	Slowdown() error
	GetSpeed() int
}

type car struct {
	engineState    bool
	handbrakeState bool
	speed          int
}

func (c *car) StartEngine() {
	c.engineState = true
}

func (c *car) StopEngine() (err error) {
	if c.speed == 0 {
		c.engineState = false
		return
	} else {
		err = errors.New("car is moving")
		return
	}
}

func (c *car) TurnOnHandbrake() (err error) {
	if (c.speed == 0) && (!c.engineState) {
		c.handbrakeState = true
	} else {
		err = errors.New("stop the car and turn off engine")
	}
	return
}

func (c *car) TurnOffHandbrake() {
	c.handbrakeState = false
}

func (c *car) StartMove() (err error) {
	if c.engineState && (!c.handbrakeState) {
		c.speed = 10
		return
	} else {
		err = errors.New("can't start moving")
		return
	}
}

func (c *car) Accelerate() (err error) {
	if c.engineState {
		c.speed += 10
	} else {
		err = errors.New("car can't moving")
	}
	return
}

func (c *car) Slowdown() (err error) {
	if (c.speed > 0) && (c.engineState) {
		c.speed -= 10
		return
	} else {
		err = errors.New("speed is zero or car can't moving")
		return
	}
}

func (c *car) GetSpeed() (speed int) {
	return c.speed
}

// NewCar  конструктор нового автомобиля
func NewCar() Car {
	return &car{
		engineState:    false,
		handbrakeState: true,
		speed:          0,
	}
}
