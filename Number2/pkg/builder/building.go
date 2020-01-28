package builder

// Building interface of buildings
type Building interface {
	CalculateDefence(map[string]int) int
	GetDoorType() string
	GetWallType() string
}

type building struct {
	wallType string
	doorType string
}

func (b *building) CalculateDefence(defenceMap map[string]int) int {
	return defenceMap[b.doorType]+defenceMap[b.wallType]
}

func (b *building) GetDoorType() string {
	return b.doorType
}

func (b *building) GetWallType() string {
	return b.wallType
}