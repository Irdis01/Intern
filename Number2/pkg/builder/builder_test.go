package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	builderTest = "builder test"
)

func TestBuilder(t *testing.T) {
	cases := []struct {
		wantBuildingWall string
		wantBuildingDoor string
	}{
		{"stone", "heavy door"},
		{"wood", "light door"},
		{"cloth", "none"},
	}
	builder := NewBuilder()
	castle := builder.NewCastle()
	house := builder.NewHouse()
	tent := builder.NewTent()
	testBuildings := []Building{castle, house, tent}
	t.Run(builderTest, func(t *testing.T) {
		for i, val := range testBuildings {
			assert.Equal(t, val.GetDoorType(), cases[i].wantBuildingDoor)
			assert.Equal(t, val.GetWallType(), cases[i].wantBuildingWall)
		}
	})
}
