package doorShop

type door interface {
	GetName() (name string)
}

// DoorShop интерфейс магазина дверей
type DoorShop interface {
	CheckDoorAvailable(doorName string) (result bool)
	AddDoor(doorInterface door)
}

type doorShop struct {
	doorSlice []door
}

// CheckDoorAvailable проверка наналичие двери в магазине
func (d *doorShop) CheckDoorAvailable(doorName string) (result bool) {
	for i := 0; i < len(d.doorSlice); i++ {
		if d.doorSlice[i].GetName() == doorName {
			return true
		}
	}
	return false
}

// AddDoor добавление двери в магазин
func (d *doorShop) AddDoor(newDoor door) {
	if d.doorSlice == nil {
		d.doorSlice = make([]door, 1)
		d.doorSlice[0] = newDoor
	} else {
		d.doorSlice = append(d.doorSlice, newDoor)
	}
}

// NewDoorShop фабрика магазина дверей
func NewDoorShop() DoorShop {
	return &doorShop{}
}
