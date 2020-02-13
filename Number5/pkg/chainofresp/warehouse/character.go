package warehouse

import (
	"fmt"
	"github.com/Irdis01/Intern/Number5/pkg/models"
)

// Warehouse warehouse interface
type Warehouse interface {
	Handle(waybill models.Waybill) (product *models.Product, err error)
}

type warehouse struct {
	productBase []*models.Product
}

func (w *warehouse) Handle(waybill models.Waybill) (product *models.Product, err error) {
	for i := 0; i < len(w.productBase); i++ {
		if w.productBase[i].Id == waybill.ProductId {
			if w.productBase[i].Name != waybill.ProductName {
				return nil, fmt.Errorf("wrong product name with right id")
			}
			defer func() {
				if i != (len(w.productBase) - 1) {
					w.productBase = append(w.productBase[:i], w.productBase[i+1:]...)
				} else {
					w.productBase = w.productBase[:i]
				}
			}()
			return w.productBase[i], nil
		}
	}
	return nil, fmt.Errorf("product not found")
}

// NewWarehouse warehouse constructor
func NewWarehouse() Warehouse {
	return &warehouse{}
}
