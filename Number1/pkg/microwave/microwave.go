package microwave

import (
	"github.com/Irdis01/Intern/Number1/pkg/models"
)

type Microwave interface {
	StartHeat(temp int)
}

type microwave struct {
	foodInMicrowave models.Food
}

func (m *microwave) StartHeat(temp int) {
	m.foodInMicrowave.Temp = temp
}

func NewMicrowave(food models.Food) Microwave {
	return &microwave{
		foodInMicrowave: food,
	}
}
