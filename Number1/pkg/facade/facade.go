package facade

type productInterface interface {
	GetName() (name string)
	GetType() (name string)
}

type doorShop interface {
	SellDoor(name string) bool
}

type paintShop interface {
	SellPaint(name string) bool
}

//Manager интерфейс фасада
type Manager interface {
	Bye(product productInterface) bool
}

type manager struct {
	newDoorShop  doorShop
	newPaintShop paintShop
}

//Bye функция, определяющая есть ли товар в магазине
func (m *manager) Bye(product productInterface) bool {
	switch product.GetType() {
	case "door":
		{
			res := m.newDoorShop.SellDoor(product.GetName())
			return res
		}
	case "paint":
		{
			res := m.newPaintShop.SellPaint(product.GetName())
			return res
		}
	}
	return false
}

//NewManager фабрика фасада
func NewManager(newDoorShop doorShop, newPaintShop paintShop) Manager {
	return &manager{newDoorShop: newDoorShop, newPaintShop: newPaintShop}
}
