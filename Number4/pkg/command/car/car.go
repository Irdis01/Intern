package car

import "errors"

type car struct {
	turnSignalState bool
	engineState     bool
	handbrakeState  bool
	speed           int
}

func (c *car) TurnOnTurnSignal() {
	c.turnSignalState = true
}

func (c *car) TurnOfTurnSignal() {
	c.turnSignalState = false
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

func (c *car) TurnOnHandbrake() {
	c.turnSignalState = true
}

func (c *car) TurnOffHandbrake() {
	c.turnSignalState = false
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

func (c *car) Accelerate() {
	c.speed += 10
}

func (c *car) Slowdown() (err error) {
	if c.speed > 0 {
		c.speed -= 10
		return
	} else {
		err = errors.New("speed is zero")
		return
	}
}
