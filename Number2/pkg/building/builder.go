package building

type builder struct {
	wallType string
	doorType string
}

//Builder интерфейс строителя
type Builder interface {
	NewCastle() Building
	NewHouse() Building
	NewTent() Building
}

func (b *builder) NewCastle() Building {
	return b.setDoor("heavy door").setWall("stone").build()
}

func (b *builder) NewHouse() Building {
	return b.setDoor("light door").setWall("wood").build()
}

func (b *builder) NewTent() Building {
	return b.setDoor("cloth door").setWall("cloth").build()
}

func (b *builder) setWall(wallType string) *builder {
	b.wallType = wallType
	return b
}

func (b *builder) setDoor(doorType string) *builder {
	b.doorType = doorType
	return b
}

func (b *builder) build() Building {
	return &building{
		wallType: b.wallType,
		doorType: b.doorType,
	}
}

//NewBuilder конструктор нового строителя
func NewBuilder() Builder {
	return &builder{}
}
