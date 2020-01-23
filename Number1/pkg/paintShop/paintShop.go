package paintShop

type paint interface {
	GetName() (name string)
}

// PaintShop Интерфейс магазина красок
type PaintShop interface {
	CheckPaintAvailable(paint string) (result bool)
	AddPaint(newPaint paint)
}

type paintShop struct {
	paintSlice []paint
}

//C heckPaintAvailable проверка наналичие краскив магазине
func (p *paintShop) CheckPaintAvailable(paintName string) (result bool) {
	for i := 0; i < len(p.paintSlice); i++ {
		if p.paintSlice[i].GetName() == paintName {
			return true
		}
	}
	return false
}

// AddPaint добавление красок в магазин
func (p *paintShop) AddPaint(newPaint paint) {
	if p.paintSlice == nil {
		p.paintSlice = make([]paint, 1)
		p.paintSlice[0] = newPaint
	} else {
		p.paintSlice = append(p.paintSlice, newPaint)
	}
}

// NewPaintShop фабрика магазина красок
func NewPaintShop() PaintShop {
	return &paintShop{}
}
