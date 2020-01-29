package warehouse

import "github.com/Irdis01/Intern/Number5/pkg/models"

// Warehouse интерфейс, описывающий работу сервиса
type Warehouse interface {
	Handle(login string, msg *[]byte) (err error)
}

type warehouse struct{
	productBase []models.Product
}

func (s *warehouse) Handle() (product models.Product, err error) {
	return
}

// NewService конструктор новго сервиса
func NewService() Warehouse {
	return &service{}
}
