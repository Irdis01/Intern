package facade

type productInterface interface {
	GetName() (name string)
}

type doorShop interface {
	CheckDoorAvailable(name string) (result bool)
}

type paintShop interface {
	CheckPaintAvailable(name string) (result bool)
}

// Manager интерфейс фасада
type Manager interface {
	Check(product productInterface) (result bool)
}

type manager struct {
	newDoorShop  doorShop
	newPaintShop paintShop
}

// Check функция, определяющая есть ли товар в магазине
func (m *manager) Check(product productInterface) bool {
	doorShopResult := m.newDoorShop.CheckDoorAvailable(product.GetName())
	paintShoppResult := m.newPaintShop.CheckPaintAvailable(product.GetName())
	return doorShopResult || paintShoppResult
}

// NewManager фабрика фасада
func NewManager(newDoorShop doorShop, newPaintShop paintShop) Manager {
	return &manager{
		newDoorShop:  newDoorShop,
		newPaintShop: newPaintShop,
	}
}
