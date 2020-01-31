package builder

type building interface {
	SetWall(wallType string)
	SetDoor(doorType string)
}

// Builder interface of builder
type Builder interface {
	NewCastle(castle building)
	NewHouse(house building)
	NewTent(tent building)
}

type builder struct{}

// NewCastle return new castle
func (b *builder) NewCastle(castle building) {
	castle.SetDoor("heavy door")
	castle.SetWall("stone")
}

// NewHouse return new house
func (b *builder) NewHouse(house building) {
	house.SetDoor("light door")
	house.SetWall("wood")
}

//NewTent return new ten
func (b *builder) NewTent(tent building) {
	tent.SetWall("cloth")
}

// NewBuilder return new builder
func NewBuilder() Builder {
	return &builder{}
}
