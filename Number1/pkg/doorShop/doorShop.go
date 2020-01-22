package doorShop

type door interface {
	GetName() (name string)
	GetType() (name string)
}

//DoorShop интерфейс магазина дверей
type DoorShop interface {
	SellDoor(doorName string) bool
	AddDoor(doorInterface door)
}

type doorShop struct {
	doorSlice []door
}

func (d *doorShop) SellDoor(doorName string) bool {
	for i := 0; i < len(d.doorSlice); i++ {
		if d.doorSlice[i].GetName() == doorName {
			return true
		}
	}
	return false
}

func (d *doorShop) AddDoor(newDoor door) {
	d.doorSlice = append(d.doorSlice, newDoor)
}

//NewDoorShop фабрика магазина дверей
func NewDoorShop() DoorShop {
	return &doorShop{
		doorSlice: make([]door, 0),
	}
}
