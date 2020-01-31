package building

import (
	"fmt"
)

// Building interface of buildings
type Building interface {
	Calculate(map[string]int) (defenceRating int, err error)
	SetWall(wallType string)
	SetDoor(doorType string)
}

type building struct {
	wallType string
	doorType string
}

func (b *building) SetWall(wallType string)  {
	b.wallType = wallType
}

func (b *building) SetDoor(doorType string)  {
	b.doorType = doorType
}

func (b *building) Calculate(defenceMap map[string]int) (defenceRating int, err error) {
	if doorDefence, ok := defenceMap[b.doorType]; !ok {
		err = fmt.Errorf("material not found in defence map")
		return
	} else {
		defenceRating += doorDefence
	}
	if wallDefence, ok := defenceMap[b.wallType]; !ok {
		err = fmt.Errorf("material not found in defence map")
		return
	} else {
		defenceRating += wallDefence
	}
	return
}

func NewBuilding() Building {
	return &building {
		wallType:"none",
		doorType:"none",
	}
}