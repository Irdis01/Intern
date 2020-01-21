package facade

import (
	"github.com/Irdis01/Intern/Number1/pkg/doorShop"
	"github.com/Irdis01/Intern/Number1/pkg/paintShop"
	"github.com/Irdis01/Intern/Number1/pkg/products"
)

type Manager interface {
	Bye(product ...products.Product)
}

type manager struct {
	newDoorShop  doorShop.DoorShop
	newPaintShop paintShop.PaintShop
}

func (m *manager) Bye(product []products.Product) []bool {
	resArray := make([]bool, 0)
	for i := 0; i < len(product); i++ {
		switch product[i].GetType() {
		case "door":
			{
				_, res := m.newDoorShop.SellDoor(product[i])
				//fmt.Println(product[i].GetName() + " " + strconv.FormatBool(res))
				resArray = append(resArray, res)
			}
		case "paint":
			{
				_, res := m.newPaintShop.SellPaint(product[i])
				//fmt.Println(product[i].GetName() + " " + strconv.FormatBool(res))
				resArray = append(resArray, res)
			}
		}
	}
	return resArray
}

func InitManager() *manager {
	var newManager manager
	newManager.newDoorShop = doorShop.NewDoorShop()
	newManager.newPaintShop = paintShop.NewpaintShop()
	return &newManager
}

// func (m *manager) Add(product ...products.Product) {
// 	for i := 0; i < len(product); i++ {
// 		switch product[i].GetType() {
// 		case "door":
// 			{
// 				m.newDoorShop.AddDoor(product[i])
// 			}
// 		case "paint":
// 			{
// 				m.newPaintShop.AddPaint(product[i])
// 			}
// 		}
// 	}
// }
