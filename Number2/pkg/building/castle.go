package building

type castle struct {
	wallType string
	doorType string
}

func (c *castle) GetWall() string {
	return c.wallType
}

func (c *castle) GetDoor() string {
	return c.doorType
}
