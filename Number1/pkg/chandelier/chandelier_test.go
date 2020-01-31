package chandelier

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	addTestSuccess     = "Add test success"
	addTestError       = "Add test error"
	removeTestSuccess  = "Remove test success"
	removeTestError    = "Remove test error"
	turnOnTestSuccess  = "Turn On test"
	turnOnTestError    = "Turn On test"
	turnOffTestSuccess = "Turn On test"
	turnOffTestError   = "Turn On test"

	addError         = "chandelier has lamp"
	removeError      = "chandelier has no lamp"
	turnOnWorksError = "lamp is already works"
	turnOnBurnError  = "lamp is burn out"
	furnOffError     = "lamp is shutdown"
)

func TestChandelier_AddSuccess(t *testing.T) {
	testChandelier := NewChandelier()
	lamps := &models.Lamp{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	}
	expect := false
	t.Run(addTestSuccess, func(t *testing.T) {
		res, err := testChandelier.Add(lamps)
		assert.NoError(t, err, "unexpecting error")
		assert.Equal(t, expect, res)
	})
}

func TestChandelier_AddError(t *testing.T) {
	testChandelier := NewChandelier()
	lamps := &models.Lamp{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	}
	testChandelier.Add(lamps)
	expect := errors.New(addError)
	t.Run(addTestError, func(t *testing.T) {
		_, err := testChandelier.Add(lamps)
		assert.Equal(t, expect, err)
	})
}

func TestChandelier_RemoveSuccess(t *testing.T) {
	testChandelier := NewChandelier()
	lamps := &models.Lamp{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	}
	expect := *lamps
	testChandelier.Add(lamps)
	t.Run(removeTestSuccess, func(t *testing.T) {
		res, err := testChandelier.Remove()
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestChandelier_RemoveError(t *testing.T) {
	testChandelier := NewChandelier()
	expect := errors.New(removeError)
	t.Run(removeTestError, func(t *testing.T) {
		_, err := testChandelier.Remove()
		assert.Equal(t, expect, err)
	})
}

func TestChandelier_TurnOnSuccess(t *testing.T) {
	//testChandelier:=NewChandelier()
	lamps := &models.Lamp{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	}
	expect := true
	testChandelier := NewChandelier()
	testChandelier.Add(lamps)
	t.Run(turnOnTestSuccess, func(t *testing.T) {
		res, err := testChandelier.TurnOn()
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestChandelier_TurnOnError(t *testing.T) {
	//testChandelier:=NewChandelier()
	lamps := []*models.Lamp{
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
			err:    errors.New(turnOnWorksError),
		},
		{
			result: false,
			err:    errors.New(turnOnBurnError),
		},
		{
			result: false,
			err:    errors.New(turnOnBurnError),
		},
	}
	t.Run(turnOnTestError, func(t *testing.T) {
		for i, _ := range expect {
			testChandelier := NewChandelier()
			testChandelier.Add(lamps[i])
			res, err := testChandelier.TurnOn()
			assert.Equal(t, expect[i].result, res)
			assert.Equal(t, expect[i].err, err)
		}
	})
}

func TestChandelier_TurnOffSuccess(t *testing.T) {
	//testChandelier:=NewChandeler()
	lamps := &models.Lamp{
		LightState:   true,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	}
	expect := false
	testChandelier := NewChandelier()
	testChandelier.Add(lamps)
	t.Run(turnOffTestSuccess, func(t *testing.T) {
		res, err := testChandelier.TurnOff()
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestChandelier_TurnOffError(t *testing.T) {
	//testChandelier:=NewChandelier()
	lamps := &models.Lamp{
		LightState:   false,
		IsBurnOut:    false,
		MaxWorkCycle: 10,
		WorkCycle:    0,
	}
	expect := errors.New("lamp is shutdown")
	testChandelier := NewChandelier()
	testChandelier.Add(lamps)
	t.Run(turnOffTestSuccess, func(t *testing.T) {
		_, err := testChandelier.TurnOff()
		assert.Equal(t, expect, err)
	})
}
