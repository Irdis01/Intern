package builder

// Builder interface of builder
type Builder interface {
	NewCastle() Building
	NewHouse() Building
	NewTent() Building
}

type builder struct {
	wallType string
	doorType string
}

// NewCastle return new castle
func (b *builder) NewCastle() Building {
	return b.setDoor("heavy door").setWall("stone").build()
}

// NewHouse return new house
func (b *builder) NewHouse() Building {
	return b.setDoor("light door").setWall("wood").build()
}

//NewTent return new ten
func (b *builder) NewTent() Building {
	return b.setWall("cloth").build()
}

func (b *builder) drop() {
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
	defer b.drop()
	return &building{
		wallType: b.wallType,
		doorType: b.doorType,
	}
}

// NewBuilder return new builder
func NewBuilder() Builder {
	return &builder{
		wallType: "none",
		doorType: "none",
	}
}
