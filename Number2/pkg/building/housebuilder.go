package building

//HouseBuilder нтерфейс строителя дома
type HouseBuilder interface {
	SetWall() HouseBuilder
	SetDoor() HouseBuilder
	GetBuilding() Building
}

type houseBuilder struct {
	wallType string
	doorType string
}

//SetWall установка вида стен
func (h *houseBuilder) SetWall() HouseBuilder {
	h.wallType = "wood"
	return h
}

//SetDoor установка вида двери
func (h *houseBuilder) SetDoor() HouseBuilder {
	h.doorType = "light door"
	return h
}

//GetBuilding постройка дома
func (h *houseBuilder) GetBuilding() Building {
	return &house{
		wallType: h.wallType,
		doorType: h.doorType,
	}
}

//NewHouseBuilder создаёт нового строителя дома
func NewHouseBuilder() HouseBuilder {
	return &houseBuilder{}
}
