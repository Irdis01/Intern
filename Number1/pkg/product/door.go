package product

type door struct {
	name string
}

func (d *door) GetName() (name string) {
	name = d.name
	return
}

func (d *door) GetType() (name string) {
	name = "door"
	return
}

//NewDoor фабрика дверей
func NewDoor(name string) Product {
	return &door{
		name: name,
	}
}
