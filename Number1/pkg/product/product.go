package product

// Product интерфейс продуктов
type Product interface {
	GetName() (name string)
}

type product struct {
	name string
}

func (p *product) GetName() (name string) {
	name = p.name
	return name
}

// NewProduct создаёт новый продукт}
func NewProduct(name string) Product {
	return &product{
		name: name,
	}
}
