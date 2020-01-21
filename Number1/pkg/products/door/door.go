package door

type ProductDoor interface {
	GetName() (name string)
	GetType() (name string)
}

type door struct {
	name string
}

func (d *door) GetName() (name string) {
	name = d.name
	return name
}

func (d *door) GetType() (name string) {
	name = "door"
	return
}

func NewDoor(name string) *door {
	newDoor := door{name: name}
	return &newDoor
}
