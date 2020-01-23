package main

import (
	"fmt"

	doorShopPkg "github.com/Irdis01/Intern/Number1/pkg/doorShop"
	"github.com/Irdis01/Intern/Number1/pkg/facade"
	paintShopPkg "github.com/Irdis01/Intern/Number1/pkg/paintShop"
	"github.com/Irdis01/Intern/Number1/pkg/product"
)

func main() {

	productArray := make([]product.Product, 0)
	productArray = append(productArray, product.NewProduct("Red"))
	productArray = append(productArray, product.NewProduct("Blue"))
	productArray = append(productArray, product.NewProduct("Green"))
	productArray = append(productArray, product.NewProduct("Wooden"))
	productArray = append(productArray, product.NewProduct("Steel"))
	productArray = append(productArray, product.NewProduct("Plastic"))

	doorShop := doorShopPkg.NewDoorShop()
	paintShop := paintShopPkg.NewPaintShop()
	doorShop.AddDoor(productArray[3])
	doorShop.AddDoor(productArray[4])
	paintShop.AddPaint(productArray[0])
	paintShop.AddPaint(productArray[1])

	manager := facade.NewManager(doorShop, paintShop)
	res := make([]bool, 0)
	for i := 0; i < len(productArray); i++ {
		res = append(res, manager.Check(productArray[i]))
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
