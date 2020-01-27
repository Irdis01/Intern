package area

import (
	"math"
)

type figure interface {
	GetSides() []float64
}

// Visiter интерфейс посетителя
type Visiter interface {
	VisitTriangle() (area interface{})
	VisitRectangle() (area interface{})
}

type areaCalculator struct {
	calcFigure figure
}

func (a *areaCalculator) VisitTriangle() (area interface{}) {
	var (
		triangleSides    []float64
		trianglePerimetr float64
	)
	triangleSides = a.calcFigure.GetSides()
	for _, val := range triangleSides {
		trianglePerimetr += val
	}
	trianglePerimetr = trianglePerimetr / 2
	area = math.Sqrt(trianglePerimetr * (trianglePerimetr - triangleSides[0]) * (trianglePerimetr - triangleSides[1]) * (trianglePerimetr - triangleSides[2]))
	return
}

func (a *areaCalculator) VisitRectangle() (area interface{}) {
	var rectangleSides []float64
	rectangleSides = a.calcFigure.GetSides()
	area = rectangleSides[0] * rectangleSides[1]
	return
}

// NewAreaCalculator конструктор посетителя, для расчёта площади площади
func NewAreaCalculator(calcFigure figure) Visiter {
	return &areaCalculator{
		calcFigure: calcFigure,
	}
}
