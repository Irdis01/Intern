package paintShop

type paint interface {
	GetName() (name string)
	GetType() (name string)
}

//PaintShop Интерфейс магазина красок
type PaintShop interface {
	CheckPaintAvailable(paint string) (result bool)
	AddPaint(newPaint paint)
}

type paintShop struct {
	paintSlice []paint
}

//CheckPaintAvailable проверка наналичие краскив магазине
func (p *paintShop) CheckPaintAvailable(paintName string) (result bool) {
	for i := 0; i < len(p.paintSlice); i++ {
		if p.paintSlice[i].GetName() == paintName {
			result = true
			return
		}
	}
	result = false
	return
}

//AddPaint добавление красок в магазин
func (p *paintShop) AddPaint(newPaint paint) {
	p.paintSlice = append(p.paintSlice, newPaint)
}

//NewpaintShop фабрика магазина красок
func NewpaintShop() PaintShop {
	return &paintShop{
		paintSlice: make([]paint, 0),
	}
}
