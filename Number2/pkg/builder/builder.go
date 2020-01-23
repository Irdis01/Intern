package building

type builder struct {
	wallType string
	doorType string
}

// NewCastle возвращает новый замок
func NewCastle() Building {
	builder := newBuilder()
	return builder.setDoor("heavy door").setWall("stone").build()
}

// NewHouse возвращает новый дом
func NewHouse() Building {
	builder := newBuilder()
	return builder.setDoor("light door").setWall("wood").build()
}

//NewTent возвращает новую палатку
func NewTent() Building {
	builder := newBuilder()
	return builder.setWall("cloth").build()
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

// NewBuilder конструктор нового строителя
func newBuilder() *builder {
	return &builder{
		wallType: "none",
		doorType: "none",
	}
}
