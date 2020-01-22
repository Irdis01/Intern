package facade

import (
	"testing"

	doorShopPkg "github.com/Irdis01/Intern/Number1/pkg/doorShop"
	paintShopPkg "github.com/Irdis01/Intern/Number1/pkg/paintShop"
	"github.com/Irdis01/Intern/Number1/pkg/product"
)

type productInterfaceTest interface {
	GetName() (name string)
	GetType() (name string)
}

func TestFacade(t *testing.T) {
	cases := []struct {
		in   productInterfaceTest
		want bool
	}{
		{product.NewPaint("Red"), true},
		{product.NewPaint("Blue"), true},
		{product.NewPaint("Green"), false},
		{product.NewDoor("Wooden"), true},
		{product.NewDoor("Steel"), true},
		{product.NewDoor("Plastic"), false},
	}

	doorShop := doorShopPkg.NewDoorShop()
	paintShop := paintShopPkg.NewpaintShop()

	doorShop.AddDoor(product.NewDoor("Wooden"))
	doorShop.AddDoor(product.NewDoor("Steel"))
	paintShop.AddPaint(product.NewPaint("Red"))
	paintShop.AddPaint(product.NewPaint("Blue"))
	manager := NewManager(doorShop, paintShop)
	for _, val := range cases {
		res := manager.Check(val.in)
		if res != val.want {
			t.Error()
		}
	}
}
