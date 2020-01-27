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
	town := make([]building.Building, 1) //город, состоящий из замка, 4 зданий и 10 палаток
	town[0] = builder.NewCastle()
	for i := 0; i < 4; i++ {
		town = append(town, builder.NewHouse())
	}
	for i := 0; i < 10; i++ {
		town = append(town, builder.NewTent())
	}
	fmt.Println(myBuilding.GetWall())
}
