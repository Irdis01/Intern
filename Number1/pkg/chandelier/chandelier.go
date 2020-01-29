package chandelier

import (
	"errors"
	"fmt"
	"github.com/Irdis01/Intern/Number1/pkg/models"
)

// Chandelier interface chandelier
type Chandelier interface {
	TurnOn() (lampState bool, err error)
	TurnOff() (lampState bool, err error)
	Add(lamp *models.Lamp) (lampIsBurn bool, err error)
	Remove() (lamp *models.Lamp, err error)
}

type chandelier struct {
	lamp *models.Lamp
}

func (c *chandelier) TurnOn() (lampState bool, err error) {
	if c.lamp.LightState {
		lampState = c.lamp.LightState
		err = fmt.Errorf("lamp is already works")
		return
	}
	c.lamp.WorkCycle++
	if c.lamp.WorkCycle == c.lamp.MaxWorkCycle {
		c.lamp.IsBurnOut = true
	}
	if c.lamp.IsBurnOut {
		err = fmt.Error("lamp is burn out")
		return
	}
	c.lamp.LightState = true
	lampState = c.lamp.LightState
	return

}

func (c *chandelier) TurnOff() (lampState bool, err error) {
	if !c.lamp.LightState {
		lampState = c.lamp.LightState
		err = fmt.Errorf("lamp is shutdown")
		return
	}
	c.lamp.LightState = false
	lampState = c.lamp.LightState
	return
}

func (c *chandelier) Add(lamp *models.Lamp) (lampIsBurn bool, err error) {
	if c.lamp != nil {
		err = fmt.Errorf("chandelier has lamp")
		lampIsBurn = c.lamp.IsBurnOut
		return
	}
	c.lamp = lamp
	return
}

func (c *chandelier) Remove() (lamp *models.Lamp, err error) {
	if c.lamp == nil {
		err = fmt.Errorf("chandelier has no lamp")
		return
	}
	lamp = c.lamp
	c.lamp = nil
	return
}

// NewLight return new chandelier
func NewChandelier() Chandelier {
	return &chandelier{}
}
