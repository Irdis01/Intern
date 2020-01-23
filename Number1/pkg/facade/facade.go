package facade

type productInterface interface {
	GetName() (name string)
	GetType() (name string)
}

type doorShop interface {
	CheckDoorAvailable(name string) (result bool)
}

type paintShop interface {
	CheckPaintAvailable(name string) (result bool)
}

//Manager интерфейс фасада
type Manager interface {
	Check(product productInterface) (result bool)
}

type manager struct {
	newDoorShop  doorShop
	newPaintShop paintShop
}

//Check функция, определяющая есть ли товар в магазине
func (m *manager) Check(product productInterface) bool {
	res1 := m.newDoorShop.CheckDoorAvailable(product.GetName())
	res2 := m.newPaintShop.CheckPaintAvailable(product.GetName())
	return res1 || res2
}

//NewManager фабрика фасада
func NewManager(newDoorShop doorShop, newPaintShop paintShop) Manager {
	return &manager{
		newDoorShop:  newDoorShop,
		newPaintShop: newPaintShop,
	}
}
