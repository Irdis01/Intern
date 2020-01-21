package doorShop

import "github.com/Irdis01/Intern/Number1/pkg/products/door"

type DoorShop interface {
	SellDoor(door door.ProductDoor) (door.ProductDoor, bool)
	AddDoor(door door.ProductDoor)
}

type doorShop struct {
	doorSlice []door.ProductDoor
}

func (d *doorShop) SellDoor(door door.ProductDoor) (door.ProductDoor, bool) {
	for i := 0; i < len(d.doorSlice); i++ {
		if d.doorSlice[i].GetName() == door.GetName() {
			defer func() {
				if (i + 1) < len(d.doorSlice) {
					d.doorSlice = append(d.doorSlice[:i], d.doorSlice[i+1:]...)
				} else {
					d.doorSlice = d.doorSlice[:i]
				}
			}()
			return d.doorSlice[i], true
		}
	}
	return nil, false
}

func (d *doorShop) AddDoor(door door.ProductDoor) {
	d.doorSlice = append(d.doorSlice, door)
}

func NewDoorShop() DoorShop {
	var newDoorShop doorShop
	newDoorShop.doorSlice = make([]door.ProductDoor, 0)
	//добавляю две двери
	newDoorShop.doorSlice = append(newDoorShop.doorSlice, door.NewDoor("Wooden"))
	newDoorShop.doorSlice = append(newDoorShop.doorSlice, door.NewDoor("Steel"))
	return &newDoorShop
}
