package figure

type visitor interface {
	VisitTriangle(Figure) interface{}
	VisitRectangle(Figure) interface{}
}

// Figure all figure interface
type Figure interface {
	GetSides() []float64
	Visit(visitor) interface{}
}
