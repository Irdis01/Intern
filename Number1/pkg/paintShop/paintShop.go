package paintShop

import "github.com/Irdis01/Intern/Number1/pkg/products/paint"

type PaintShop interface {
	SellPaint(paint paint.ProductPaint) (paint.ProductPaint, bool)
	AddPaint(paint paint.ProductPaint)
}

type paintShop struct {
	paintSlice []paint.ProductPaint
}

func (p *paintShop) SellPaint(paint paint.ProductPaint) (paint.ProductPaint, bool) {
	for i := 0; i < len(p.paintSlice); i++ {
		if p.paintSlice[i].GetName() == paint.GetName() {
			defer func() {
				if (i + 1) < len(p.paintSlice) {
					p.paintSlice = append(p.paintSlice[:i], p.paintSlice[i+1:]...)
				} else {
					p.paintSlice = p.paintSlice[:i]
				}
			}()
			return p.paintSlice[i], true
		}
	}
	return nil, false
}

func (p *paintShop) AddPaint(paint paint.ProductPaint) {
	p.paintSlice = append(p.paintSlice, paint)
}

func NewpaintShop() *paintShop {
	var newpaintShop paintShop
	newpaintShop.paintSlice = make([]paint.ProductPaint, 0)
	//добавляю две тетсовые краски
	newpaintShop.paintSlice = append(newpaintShop.paintSlice, paint.NewPaint("Red"))
	newpaintShop.paintSlice = append(newpaintShop.paintSlice, paint.NewPaint("Blue"))
	return &newpaintShop
}
