package figure

type triangle struct {
	firstSide     float64
	secondSide    float64
	thirdSide     float64
	figureVisitor visitor
}

func (t *triangle) GetSides() []float64 {
	return []float64{
		t.firstSide,
		t.secondSide,
		t.thirdSide,
	}
}

func (t *triangle) Visit(v visitor) interface{} {
	return v.VisitTriangle(t)
}

// NewTriangle создаёт новый треугольник
func NewTriangle(firstSide float64, secondSide float64, thirdSide float64) Figure {
	return &triangle{
		firstSide:  firstSide,
		secondSide: secondSide,
		thirdSide:  thirdSide,
	}
}
