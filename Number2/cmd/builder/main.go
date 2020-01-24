package main

import (
	"fmt"

	building "github.com/Irdis01/Intern/Number2/pkg/builder"
)

func main() {
	var (
		myBuilding building.Building
	)
	builder := building.NewBuilder()
	buildingType := "tent"
	switch buildingType {
	case "castle":
		{
			myBuilding = builder.NewCastle()
		}
	case "house":
		{
			myBuilding = builder.NewHouse()
		}
	case "tent":
		{
			myBuilding = builder.NewTent()
		}
	}

	fmt.Println(myBuilding.GetWall())

}
