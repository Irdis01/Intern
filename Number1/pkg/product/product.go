package product

//Product интерфейс продуктов
type Product interface {
	GetName() (name string)
	GetType() (productType string)
}
