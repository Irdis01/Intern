package main

import (
	"fmt"

	"github.com/Irdis01/Intern/Number2/pkg/building"
)

func main() {
	var (
		myBuilding building.Building
		myBuilder  building.Builder
	)
	buildingType := "tent"
	myBuilder = building.NewBuilder()
	switch buildingType {
	case "castle":
		{
			myBuilding = myBuilder.NewCastle()
		}
	case "house":
		{
			myBuilding = myBuilder.NewHouse()
		}
	case "tent":
		{
			myBuilding = myBuilder.NewTent()
		}
	}

	fmt.Println(myBuilding.GetWall())

}
