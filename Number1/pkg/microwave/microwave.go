package microwave

import (
	"fmt"
	"github.com/Irdis01/Intern/Number1/pkg/models"
)

// Microwave microwave interface
type Microwave interface {
	Heat(temp int) (foodTemp int, err error)
	Add(food *models.Food) (err error)
	Get() (foodInMicrowave *models.Food, err error)
}

type microwave struct {
	foodInMicrowave *models.Food
}

func (m *microwave) Heat(temp int) (foodTemp int, err error) {
	if m.foodInMicrowave == nil {
		err = fmt.Errorf("no food in microwave")
		return
	}
	m.foodInMicrowave.Temp = temp
	foodTemp = m.foodInMicrowave.Temp
	return
}

func (m *microwave) Add(food *models.Food) (err error) {
	if m.foodInMicrowave != nil {
		err = fmt.Errorf("microwave is not empty")
		return
	}
	m.foodInMicrowave = food
	return
}

func (m *microwave) Get() (foodInMicrowave *models.Food, err error) {
	if m.foodInMicrowave == nil {
		err = fmt.Errorf("microwave is empty")
		return
	}
	foodInMicrowave = m.foodInMicrowave
	m.foodInMicrowave = nil
	return
}

func NewMicrowave() Microwave {
	return &microwave{}
}
