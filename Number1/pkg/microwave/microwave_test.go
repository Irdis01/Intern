package microwave

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	addTestSuccess  = "Add test success"
	addTestError    = "Add test error"
	getTestSuccess  = "Get test success"
	getTestError    = "Get test error"
	heatTestSuccess = "Heat test success"
	heatTestError   = "Heat test error"

	addError  = "microwave is not empty"
	getError  = "microwave is empty"
	heatError = "no food in microwave"
)

func TestMicrowave_AddSuccess(t *testing.T) {
	TestMicrowave := NewMicrowave()
	foods := &models.Food{
		Temp: 0,
	}
	expect := error(nil)
	t.Run(addTestSuccess, func(t *testing.T) {
		err := TestMicrowave.Add(foods)
		assert.Equal(t, expect, err)
	})
}

func TestMicrowave_AddError(t *testing.T) {
	testMicrowave := NewMicrowave()
	foods := &models.Food{
		Temp: 0,
	}
	expect := errors.New(addError)
	testMicrowave.Add(foods)
	t.Run(addTestError, func(t *testing.T) {
		err := testMicrowave.Add(foods)
		assert.Equal(t, expect, err)
	})
}

func TestMicrowave_GetSuccess(t *testing.T) {
	TestMicrowave := NewMicrowave()
	foods := &models.Food{
		Temp: 0,
	}
	expect := models.Food{}
	TestMicrowave.Add(foods)
	t.Run(getTestSuccess, func(t *testing.T) {
		res, err := TestMicrowave.Get()
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestMicrowave_GetError(t *testing.T) {
	TestMicrowave := NewMicrowave()
	expect := errors.New(getError)
	t.Run(getTestError, func(t *testing.T) {
		_, err := TestMicrowave.Get()
		assert.Equal(t, expect, err)
	})
}

func TestMicrowave_HeatSuccess(t *testing.T) {
	TestMicrowave := NewMicrowave()
	foods := &models.Food{
		Temp: 0,
	}
	expect := 80
	TestMicrowave.Add(foods)
	t.Run(heatTestSuccess, func(t *testing.T) {
		res, err := TestMicrowave.Heat(80)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestMicrowave_HeatError(t *testing.T) {
	var foods *models.Food
	expect := errors.New(heatError)
	TestMicrowave := NewMicrowave()
	TestMicrowave.Add(foods)
	t.Run(heatTestError, func(t *testing.T) {
		_, err := TestMicrowave.Heat(80)
		assert.Equal(t, expect, err)
	})
}
