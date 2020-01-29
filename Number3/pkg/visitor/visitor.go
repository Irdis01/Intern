package visitor

import (
	"math"

	"github.com/Irdis01/Intern/Number3/pkg/figure"
)

// Visitor visitor interface
type Visitor interface {
	VisitTriangle(geomtricFigure figure.Figure) (area interface{})
	VisitRectangle(geomtricFigure figure.Figure) (area interface{})
}

type areaCalculator struct{}

func (a *areaCalculator) VisitTriangle(geomtricFigure figure.Figure) (area interface{}) {
	var (
		triangleSides    []float64
		trianglePerimetr float64
	)
	triangleSides = geomtricFigure.GetSides()
	for _, val := range triangleSides {
		trianglePerimetr += val
	}
	trianglePerimetr = trianglePerimetr / 2
	area = math.Sqrt(trianglePerimetr * (trianglePerimetr - triangleSides[0]) * (trianglePerimetr - triangleSides[1]) * (trianglePerimetr - triangleSides[2]))
	return
}

func (a *areaCalculator) VisitRectangle(geomtricFigure figure.Figure) (area interface{}) {
	var rectangleSides []float64
	rectangleSides = geomtricFigure.GetSides()
	area = rectangleSides[0] * rectangleSides[1]
	return
}

// NewVisitor return bew visitor
func NewVisitor() Visitor {
	return &areaCalculator{}
}
