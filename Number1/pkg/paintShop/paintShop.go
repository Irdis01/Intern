package paintShop

type paint interface {
	GetName() (name string)
	GetType() (name string)
}

//PaintShop Интерфейс магазина красок
type PaintShop interface {
	SellPaint(paint string) bool
	AddPaint(newPaint paint)
}

type paintShop struct {
	paintSlice []paint
}

func (p *paintShop) SellPaint(paintName string) bool {
	for i := 0; i < len(p.paintSlice); i++ {
		if p.paintSlice[i].GetName() == paintName {
			return true
		}
	}
	return false
}

func (p *paintShop) AddPaint(newPaint paint) {
	p.paintSlice = append(p.paintSlice, newPaint)
}

//NewpaintShop фабрика магазина красок
func NewpaintShop() PaintShop {
	var newpaintShop paintShop
	newpaintShop.paintSlice = make([]paint, 0)
	return &newpaintShop
}
