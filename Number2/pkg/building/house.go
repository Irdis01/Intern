package building

type house struct {
	wallType string
	doorType string
}

func (h *house) GetWall() string {
	return h.wallType
}

func (h *house) GetDoor() string {
	return h.doorType
}
