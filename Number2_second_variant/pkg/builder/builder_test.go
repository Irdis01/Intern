package builder

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	newCastleTest                = "newCastleTest"
	newHouseTest                 = "newHouseTest"
	newTentTest                  = "newTentTest"
	defenceCalculatorTestSuccess = "defenceCalculator success test"
	defenceCalculatorTestError   = "defenceCalculator error test"

	defenceCalculatorMaterialNotFoundError = "material not found in defence map"
)

func TestBuilder_NewCastle(t *testing.T) {
	var defenceMap = map[string]int{
		"stone":      10,
		"heavy door": 5,
	}
	expect := 15
	builder := NewBuilder()
	castle := builder.NewCastle()
	t.Run(newCastleTest, func(t *testing.T) {
		res, err := castle.Calculate(defenceMap)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestBuilder_NewHouse(t *testing.T) {
	var defenceMap = map[string]int{
		"wood":       5,
		"light door": 2,
	}
	expect := 7
	builder := NewBuilder()
	house := builder.NewHouse()
	t.Run(newHouseTest, func(t *testing.T) {
		res, err := house.Calculate(defenceMap)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestBuilder_NewTent(t *testing.T) {
	var defenceMap = map[string]int{
		"cloth": 2,
		"none":  0,
	}
	expect := 2
	builder := NewBuilder()
	tent := builder.NewTent()
	t.Run(newTentTest, func(t *testing.T) {
		res, err := tent.Calculate(defenceMap)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestBuilding_CalculateDefenceSuccess(t *testing.T) {
	var defenceMap = map[string]int{
		"stone":      10,
		"heavy door": 5,
	}
	expect := 15
	builder := NewBuilder()
	castle := builder.NewCastle()
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
	builder := NewBuilder()
	castle := builder.NewHouse()
	t.Run(defenceCalculatorTestSuccess, func(t *testing.T) {
		_, err := castle.Calculate(defenceMap)
		assert.Equal(t, expect, err)
	})
}
