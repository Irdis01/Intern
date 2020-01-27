package builder

// Builder интерфейс строителя
type Builder interface {
	NewCastle() Building
	NewHouse() Building
	NewTent() Building
}

type builder struct {
	wallType string
	doorType string
}

// NewCastle возвращает новый замок
func (b *builder) NewCastle() Building {
	return b.setDoor("heavy door").setWall("stone").build()
}

// NewHouse возвращает новый дом
func (b *builder) NewHouse() Building {
	return b.setDoor("light door").setWall("wood").build()
}

//NewTent возвращает новую палатку
func (b *builder) NewTent() Building {
	return b.setWall("cloth").build()
}

//DropBuilder сбрасывает поля строителя
func (b *builder) dropBuilder() {
	b.setDoor("none")
	b.setWall("none")
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
	defer b.dropBuilder()
	return &building{
		wallType: b.wallType,
		doorType: b.doorType,
	}
}

// NewBuilder конструктор нового строителя
func NewBuilder() Builder {
	return &builder{
		wallType: "none",
		doorType: "none",
	}
}
