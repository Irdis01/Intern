package door

//ProductDoor интерфейс дверей
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

//NewDoor фабрика дверей
func NewDoor(name string) ProductDoor {
	newDoor := door{name: name}
	return &newDoor
}
