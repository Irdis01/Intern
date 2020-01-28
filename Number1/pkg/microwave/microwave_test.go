package microwave

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	AddTest  = "Add test"
	GetTest  = "Get test"
	HeatTest = "Heat test"
)

func TestMicrowave_Add(t *testing.T) {
	TestMicrowave := NewMicrowave()
	foods := []*models.Food{{
		0,
	},
		{
			Temp: 0,
		},
	}
	expect := []struct {
		result int
		err    error
	}{
		{
			err: nil,
		},
		{
			err: errors.New("microwave is not empty"),
		},
	}
	t.Run(AddTest, func(t *testing.T) {
		for i, _ := range expect {
			err := TestMicrowave.Add(foods[i])
			assert.Equal(t, expect[i].err, err)
		}
	})

}

func TestMicrowave_Get(t *testing.T) {
	TestMicrowave := NewMicrowave()
	foods := []*models.Food{{
		0,
	},
		{
			Temp: 0,
		},
	}
	expect := []struct {
		result *models.Food
		err    error
	}{
		{
			result: foods[0],
			err:    nil,
		},
		{
			result: nil,
			err:    errors.New("microwave is empty"),
		},
	}
	TestMicrowave.Add(foods[0])
	t.Run(GetTest, func(t *testing.T) {
		for i, _ := range expect {
			res, err := TestMicrowave.Get()
			assert.Equal(t, expect[i].err, err)
			assert.Equal(t, expect[i].result, res)
		}
	})
}

func TestMicrowave_Heat(t *testing.T) {
	foods := []*models.Food{{
		0,
	},
		nil,
	}
	expect := []struct {
		result int
		err    error
	}{
		{
			result: 80,
			err:    nil,
		},
		{
			result: 0,
			err:    errors.New("no food in microwave"),
		},
	}
	t.Run(HeatTest, func(t *testing.T) {
		for i, _ := range expect {
			TestMicrowave := NewMicrowave()
			TestMicrowave.Add(foods[i])
			res, err := TestMicrowave.Heat(80)
			assert.Equal(t, expect[i].err, err)
			assert.Equal(t, expect[i].result, res)
		}
	})
}
