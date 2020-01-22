package building

//CastleBuilder интерфейс строителя замка
type CastleBuilder interface {
	SetWall() CastleBuilder
	SetDoor() CastleBuilder
	GetBuilding() Building
}

type castleBuilder struct {
	wallType string
	doorType string
}

//SetWall установка вида стен
func (c *castleBuilder) SetWall() CastleBuilder {
	c.wallType = "stone"
	return c
}

//SetDoor установка вида двери
func (c *castleBuilder) SetDoor() CastleBuilder {
	c.doorType = "heavy door"
	return c
}

//GetBuilding постройка замка
func (c *castleBuilder) GetBuilding() Building {
	return &castle{
		wallType: c.wallType,
		doorType: c.doorType,
	}
}

//NewCastleBuilder инициализация новго строителя замка
func NewCastleBuilder() CastleBuilder {
	return &castleBuilder{}
}
