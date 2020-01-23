package building

// Building интерфейс зданий
type Building interface {
	GetWall() string
	GetDoor() string
}

type building struct {
	wallType string
	doorType string
}

func (b *building) GetWall() string {
	return b.wallType
}

func (b *building) GetDoor() string {
	return b.doorType
}
