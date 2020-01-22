package building

//TentBuilder интерфейс строителя палатки
type TentBuilder interface {
	SetWall() TentBuilder
	SetDoor() TentBuilder
	GetBuilding() Building
}

type tentBuilder struct {
	wallType string
	doorType string
}

type tent struct {
	wallType string
	doorType string
}

func (t *tent) GetWall() string {
	return t.wallType
}

func (t *tent) GetDoor() string {
	return t.doorType
}

//SetWall установка вида стен
func (t *tentBuilder) SetWall() TentBuilder {
	t.wallType = "stone"
	return t
}

//SetDoor установка вида двери
func (t *tentBuilder) SetDoor() TentBuilder {
	t.doorType = "heavy door"
	return t
}

//GetBuilding постройка палатеи
func (t *tentBuilder) GetBuilding() Building {
	return &tent{
		wallType: t.wallType,
		doorType: t.doorType,
	}
}

//NewTentBuilder инициализация новго строителя палатки
func NewTentBuilder() TentBuilder {
	return &tentBuilder{}
}
