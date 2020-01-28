package chandelier

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	AddTest     = "Add test"
	RemoveTest  = "Remove test"
	TurnOnTest  = "Turn On test"
	TurnOffTest = "Turn Off test"
)

func TestChandelier_Add(t *testing.T) {
	testChandelier := NewChandelier()
	lamps := []*models.Lamp{{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	},
		{
			LightState:   false,
			IsBurnOut:    false,
			MaxWorkCycle: 10,
			WorkCycle:    0,
		},
	}
	expect := []struct {
		result bool
		err    error
	}{
		{
			result: false,
			err:    nil,
		},
		{
			result: false,
			err:    errors.New("chandelier has lamp"),
		},
	}
	t.Run(AddTest, func(t *testing.T) {
		for i, _ := range expect {
			res, err := testChandelier.Add(lamps[i])
			assert.Equal(t, expect[i].result, res)
			assert.Equal(t, expect[i].err, err)
		}
	})
}

func TestChandelier_Remove(t *testing.T) {
	testChandelier := NewChandelier()
	lamps := []*models.Lamp{{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	},
		{
			LightState:   false,
			IsBurnOut:    false,
			MaxWorkCycle: 10,
			WorkCycle:    0,
		},
	}
	expect := []struct {
		result *models.Lamp
		err    error
	}{
		{
			result: lamps[0],
			err:    nil,
		},
		{
			result: nil,
			err:    errors.New("chandelier has no lamp"),
		},
	}
	testChandelier.Add(lamps[0])
	t.Run(RemoveTest, func(t *testing.T) {
		for i, _ := range expect {
			res, err := testChandelier.Remove()
			assert.Equal(t, expect[i].result, res)
			assert.Equal(t, expect[i].err, err)
		}
	})
}

func TestChandelier_TurnOn(t *testing.T) {
	//testChandelier:=NewChandelier()
	lamps := []*models.Lamp{{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	},
		{
			LightState:   true,
			IsBurnOut:    false,
			MaxWorkCycle: 10,
			WorkCycle:    0,
		},
		{
			LightState:   false,
			IsBurnOut:    true,
			MaxWorkCycle: 10,
			WorkCycle:    0,
		},
		{
			LightState:   false,
			IsBurnOut:    false,
			MaxWorkCycle: 1,
			WorkCycle:    0,
		},
	}
	expect := []struct {
		result bool
		err    error
	}{
		{
			result: true,
			err:    nil,
		},
		{
			result: true,
			err:    errors.New("lamp is already works"),
		},
		{
			result: false,
			err:    errors.New("lamp is burn out"),
		},
		{
			result: false,
			err:    errors.New("lamp is burn out"),
		},
	}
	t.Run(TurnOnTest, func(t *testing.T) {
		for i, _ := range expect {
			testChandelier := NewChandelier()
			testChandelier.Add(lamps[i])
			res, err := testChandelier.TurnOn()
			assert.Equal(t, expect[i].result, res)
			assert.Equal(t, expect[i].err, err)
		}
	})
}

func TestChandelier_TurnOff(t *testing.T) {
	//testChandelier:=NewChandelier()
	lamps := []*models.Lamp{{
		LightState:   true,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	},
		{
			LightState:   false,
			IsBurnOut:    false,
			MaxWorkCycle: 10,
			WorkCycle:    0,
		},
	}
	expect := []struct {
		result bool
		err    error
	}{
		{
			result: false,
			err:    nil,
		},
		{
			result: false,
			err:    errors.New("lamp is shutdown"),
		},
	}
	testChandelier := NewChandelier()
	testChandelier.Add(lamps[0])
	t.Run(TurnOffTest, func(t *testing.T) {
		for i, _ := range expect {
			testChandelier := NewChandelier()
			testChandelier.Add(lamps[i])
			res, err := testChandelier.TurnOff()
			assert.Equal(t, expect[i].result, res)
			assert.Equal(t, expect[i].err, err)
		}
	})
}
