package building

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	cases := []struct {
		wantBuildingWall string
		wantBuildingDoor string
	}{
		{"stone", "heavy door"},
		{"wood", "light door"},
		{"cloth", "cloth door"},
	}
	builder := NewBuilder()

	castle := builder.NewCastle()
	if (castle.GetWall() != cases[0].wantBuildingWall) || (castle.GetDoor() != cases[0].wantBuildingDoor) {
		t.Error()
	}
	house := builder.NewHouse()
	if (house.GetWall() != cases[1].wantBuildingWall) || (house.GetDoor() != cases[1].wantBuildingDoor) {
		t.Error()
	}
	tent := builder.NewTent()
	if (tent.GetWall() != cases[2].wantBuildingWall) || (tent.GetDoor() != cases[2].wantBuildingDoor) {
		t.Error()
	}
}
