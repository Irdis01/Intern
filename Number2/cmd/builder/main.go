package main

import (
	"fmt"

	"github.com/Irdis01/Intern/Number2/pkg/building"
)

func main() {
	var myBuilding building.Building
	builderType := "tent"
	switch builderType {
	case "castle":
		{
			builder := building.NewCastleBuilder()
			myBuilding = builder.SetDoor().SetWall().GetBuilding()
		}
	case "house":
		{
			builder := building.NewHouseBuilder()
			myBuilding = builder.SetDoor().SetWall().GetBuilding()
		}
	case "tent":
		{
			builder := building.NewTentBuilder()
			myBuilding = builder.SetDoor().SetWall().GetBuilding()
		}
	}

	fmt.Println(myBuilding.GetWall())

}
