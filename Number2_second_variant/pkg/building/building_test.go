package building

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	setDoorTest                  = "SetDoor test"
	defenceCalculatorTestSuccess = "defenceCalculator success test"
	defenceCalculatorTestError   = "defenceCalculator error test"

	defenceCalculatorMaterialNotFoundError = "material not found in defence map"
)

func TestBuilding_CalculateDefenceSuccess(t *testing.T) {
	var defenceMap = map[string]int{
		"stone":      10,
		"heavy door": 5,
	}
	expect := 15
	castle := NewBuilding()
	castle.SetDoor("heavy door")
	castle.SetWall("stone")
	t.Run(defenceCalculatorTestSuccess, func(t *testing.T) {
		res, err := castle.Calculate(defenceMap)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestBuilding_CalculateDefenceError(t *testing.T) {
	var defenceMap = map[string]int{
		"stone":      10,
		"heavy door": 5,
	}
	expect := errors.New(defenceCalculatorMaterialNotFoundError)
	testBuilding := NewBuilding()
	t.Run(defenceCalculatorTestError, func(t *testing.T) {
		_, err := testBuilding.Calculate(defenceMap)
		assert.Equal(t, expect, err)
	})
}
