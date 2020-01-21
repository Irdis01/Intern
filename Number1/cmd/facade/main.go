package main

import (
	"fmt"

	"github.com/Irdis01/Intern/Number1/pkg/facade"
	"github.com/Irdis01/Intern/Number1/pkg/products"
	"github.com/Irdis01/Intern/Number1/pkg/products/door"
	"github.com/Irdis01/Intern/Number1/pkg/products/paint"
)

func main() {
	manager := facade.InitManager()
	productArray := make([]products.Product, 0)
	productArray = append(productArray, paint.NewPaint("Red"))
	productArray = append(productArray, door.NewDoor("Wooden"))
	productArray = append(productArray, paint.NewPaint("Blue"))
	productArray = append(productArray, door.NewDoor("Steel"))
	productArray = append(productArray, paint.NewPaint("Green"))
	productArray = append(productArray, door.NewDoor("Plastic"))
	res := manager.Bye(productArray)
	test := "PASS"
	checkArray := []bool{true, true, true, true, false, false}
	for i, _ := range res {
		if res[i] != checkArray[i] {
			test = "NOT PASS"
		}
	}
	fmt.Println(test)
}
