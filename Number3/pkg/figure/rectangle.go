package figure

type rectangle struct {
	firstSide  float64
	secondSide float64
}

func (r *rectangle) GetSides() []float64 {
	return []float64{
		r.firstSide,
		r.secondSide,
	}
}

func (r *rectangle) Visit(v visitor) interface{} {
	return v.VisitRectangle(r)
}

// NewRectangle создаёт новый прямоугольник
func NewRectangle(firstSide float64, secondSide float64) Figure {
	return &rectangle{
		firstSide:  firstSide,
		secondSide: secondSide,
	}
}
