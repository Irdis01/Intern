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

//SetWall установка вида стен
func (t *tentBuilder) SetWall() TentBuilder {
	t.wallType = "cloth"
	return t
}

//SetDoor установка вида двери
func (t *tentBuilder) SetDoor() TentBuilder {
	t.doorType = "cloth door"
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
