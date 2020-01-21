package products

//Product общий интерфейс всех продуктов
type Product interface {
	GetName() (name string)
	GetType() (name string)
}
