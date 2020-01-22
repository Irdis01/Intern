package main

import (
	"testing"

	"github.com/Irdis01/Intern/Number1/pkg/doorShop"
	"github.com/Irdis01/Intern/Number1/pkg/facade"
	"github.com/Irdis01/Intern/Number1/pkg/paintShop"
	"github.com/Irdis01/Intern/Number1/pkg/products/door"
	"github.com/Irdis01/Intern/Number1/pkg/products/paint"
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
		{paint.NewPaint("Red"), true},
		{paint.NewPaint("Blue"), true},
		{paint.NewPaint("Green"), false},
		{door.NewDoor("Wooden"), true},
		{door.NewDoor("Steel"), true},
		{door.NewDoor("Plastic"), false},
	}

	newDoorShop := doorShop.NewDoorShop()
	newPaintShop := paintShop.NewpaintShop()

	newDoorShop.AddDoor(door.NewDoor("Wooden"))
	newDoorShop.AddDoor(door.NewDoor("Steel"))
	newPaintShop.AddPaint(paint.NewPaint("Red"))
	newPaintShop.AddPaint(paint.NewPaint("Blue"))
	manager := facade.NewManager(newDoorShop, newPaintShop)

	for _, val := range cases {
		res := manager.Check(val.in)
		if res != val.want {
			t.Error()
		}
	}

}
