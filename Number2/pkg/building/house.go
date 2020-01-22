package building

//HouseBuilder нтерфейс строителя дома
type HouseBuilder interface {
	SetWall() HouseBuilder
	SetDoor() HouseBuilder
	GetBuilding() Building
}

type house struct {
	wallType string
	doorType string
}

type houseBuilder struct {
	wallType string
	doorType string
}

func (h *house) GetWall() string {
	return h.wallType
}

func (h *house) GetDoor() string {
	return h.doorType
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
