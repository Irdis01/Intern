package builder

// Building интерфейс зданий
type Building interface {
	CalculateDefence(map[string]int) int
}

type building struct {
	wallType string
	doorType string
}

func (b *building) CalculateDefence(defenceMap map[string]int) int {
	return defenceMap[b.doorType]+defenceMap[b.wallType]
}
