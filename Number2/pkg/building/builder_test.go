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
	castleBuilder := NewCastleBuilder()
	houseBuilder := NewHouseBuilder()
	tentBuilder := NewTentBuilder()

	castle := castleBuilder.SetWall().SetDoor().GetBuilding()
	if (castle.GetWall() != cases[0].wantBuildingWall) || (castle.GetDoor() != cases[0].wantBuildingDoor) {
		t.Error()
	}
	house := houseBuilder.SetWall().SetDoor().GetBuilding()
	if (house.GetWall() != cases[1].wantBuildingWall) || (house.GetDoor() != cases[1].wantBuildingDoor) {
		t.Error()
	}
	tent := tentBuilder.SetWall().SetDoor().GetBuilding()
	if (tent.GetWall() != cases[2].wantBuildingWall) || (tent.GetDoor() != cases[2].wantBuildingDoor) {
		t.Error()
	}
}
