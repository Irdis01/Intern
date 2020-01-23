package main

import (
	"fmt"

	building "github.com/Irdis01/Intern/Number2/pkg/builder"
)

func main() {
	var (
		myBuilding building.Building
	)
	buildingType := "tent"
	switch buildingType {
	case "castle":
		{
			myBuilding = building.NewCastle()
		}
	case "house":
		{
			myBuilding = building.NewHouse()
		}
	case "tent":
		{
			myBuilding = building.NewTent()
		}
	}

	fmt.Println(myBuilding.GetWall())

}
