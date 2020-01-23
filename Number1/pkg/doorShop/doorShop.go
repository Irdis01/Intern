package doorShop

type door interface {
	GetName() (name string)
	GetType() (name string)
}

//DoorShop интерфейс магазина дверей
type DoorShop interface {
	CheckDoorAvailable(doorName string) (result bool)
	AddDoor(doorInterface door)
}

type doorShop struct {
	doorSlice []door
}

//CheckDoorAvailable проверка наналичие двери в магазине
func (d *doorShop) CheckDoorAvailable(doorName string) (result bool) {
	for i := 0; i < len(d.doorSlice); i++ {
		if d.doorSlice[i].GetName() == doorName {
			result = true
			return
		}
	}
	result = false
	return
}

//AddDoor добавление двери в магазин
func (d *doorShop) AddDoor(newDoor door) {
	d.doorSlice = append(d.doorSlice, newDoor)
}

//NewDoorShop фабрика магазина дверей
func NewDoorShop() DoorShop {
	return &doorShop{
		doorSlice: make([]door, 0),
	}
}
