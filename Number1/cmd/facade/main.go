package main

import (
	"fmt"

	"github.com/Irdis01/Intern/Number1/pkg/doorShop"
	"github.com/Irdis01/Intern/Number1/pkg/facade"
	"github.com/Irdis01/Intern/Number1/pkg/paintShop"
	"github.com/Irdis01/Intern/Number1/pkg/products/door"
	"github.com/Irdis01/Intern/Number1/pkg/products/paint"
)

type productInterface interface {
	GetName() (name string)
	GetType() (name string)
}

func main() {

	productArray := make([]productInterface, 0)
	productArray = append(productArray, paint.NewPaint("Red"))
	productArray = append(productArray, paint.NewPaint("Blue"))
	productArray = append(productArray, paint.NewPaint("Green"))
	productArray = append(productArray, door.NewDoor("Wooden"))
	productArray = append(productArray, door.NewDoor("Steel"))
	productArray = append(productArray, door.NewDoor("Plastic"))

	newDoorShop := doorShop.NewDoorShop()
	newPaintShop := paintShop.NewpaintShop()
	newDoorShop.AddDoor(productArray[3])
	newDoorShop.AddDoor(productArray[4])
	newPaintShop.AddPaint(productArray[0])
	newPaintShop.AddPaint(productArray[1])

	manager := facade.NewManager(newDoorShop, newPaintShop)
	res := make([]bool, 0)
	for i := 0; i < len(productArray); i++ {
		res = append(res, manager.Bye(productArray[i]))
	}
	test := "PASS"
	checkArray := []bool{true, true, false, true, true, false}
	for i, _ := range res {
		if res[i] != checkArray[i] {
			test = "NOT PASS"
		}
	}
	fmt.Println(test)
}
