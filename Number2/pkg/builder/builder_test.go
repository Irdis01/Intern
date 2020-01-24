package building

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	testBuildings := []Building{NewCastle(), NewHouse(), NewTent()}
	t.Run("Builder test", func(t *testing.T) {
		for i, val := range testBuildings {
			assert.Equal(t, val.GetDoor(), cases[i].wantBuildingDoor)
			assert.Equal(t, val.GetWall(), cases[i].wantBuildingWall)
		}
	})
}
