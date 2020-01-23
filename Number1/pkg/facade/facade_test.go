package facade

import (
	"testing"

	doorShopPkg "github.com/Irdis01/Intern/Number1/pkg/doorShop"
	paintShopPkg "github.com/Irdis01/Intern/Number1/pkg/paintShop"
	"github.com/Irdis01/Intern/Number1/pkg/product"
)

type productInterfaceTest interface {
	GetName() (name string)
}

func TestFacade(t *testing.T) {
	cases := []struct {
		in   productInterfaceTest
		want bool
	}{
		{product.NewProduct("Red"), true},
		{product.NewProduct("Blue"), true},
		{product.NewProduct("Green"), false},
		{product.NewProduct("Wooden"), true},
		{product.NewProduct("Steel"), true},
		{product.NewProduct("Plastic"), false},
	}

	doorShop := doorShopPkg.NewDoorShop()
	paintShop := paintShopPkg.NewPaintShop()

	doorShop.AddDoor(product.NewProduct("Wooden"))
	doorShop.AddDoor(product.NewProduct("Steel"))
	paintShop.AddPaint(product.NewProduct("Red"))
	paintShop.AddPaint(product.NewProduct("Blue"))
	manager := NewManager(doorShop, paintShop)

	t.Run("facadeTest", func(t *testing.T) {
		for _, val := range cases {
			res := manager.Check(val.in)
			if res != val.want {
				t.Error()
			}
		}
	})
}
