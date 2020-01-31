package main

import (
	"github.com/Irdis01/Intern/Number2/pkg/builder"
	"log"
)

func main() {
	townBuilder := builder.NewBuilder()
	town := make([]builder.Building, 1) //город, состоящий из замка, 4 зданий и 10 палаток
	town[0] = townBuilder.NewCastle()
	for i := 0; i < 4; i++ {
		town = append(town, townBuilder.NewHouse())
	}
	for i := 0; i < 10; i++ {
		town = append(town, townBuilder.NewTent())
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
		//defenceRaiting += val.Calculate(defenceMap)
	}
	log.Println(defenceRaiting)
}
