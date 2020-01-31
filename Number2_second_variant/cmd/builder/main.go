package main

import (
	"github.com/Irdis01/Intern/Number2_second_variant/pkg/builder"
	"github.com/Irdis01/Intern/Number2_second_variant/pkg/building"
	"log"
)

func main() {
	townBuilder := builder.NewBuilder()
	town := make([]building.Building, 1) //город, состоящий из замка, 4 зданий и 10 палаток
	castle := building.NewBuilding()
	townBuilder.NewCastle(castle)
	town[0]=castle
	for i := 0; i < 4; i++ {
		house := building.NewBuilding()
		townBuilder.NewHouse(house)
		town = append(town, house)
	}
	for i := 0; i < 10; i++ {
		tent := building.NewBuilding()
		townBuilder.NewTent(tent)
		town = append(town, tent)
	}
	var defenceRaiting int
	defenceMap := map[string]int{
		"stone":      10,
		"wood":       5,
		"cloth":      2,
		"heavy door": 5,
		"light door": 2,
		"none":       0,
	}
	for _, val := range town {
		rating, _ := val.Calculate(defenceMap)
		defenceRaiting += rating
	}
	log.Println(defenceRaiting)
}
