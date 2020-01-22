package building

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
