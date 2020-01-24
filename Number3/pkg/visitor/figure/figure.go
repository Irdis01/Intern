package figure

type visitor interface {
	VisitTriangle(Figure) interface{}
	VisitRectangle(Figure) interface{}
}

// Figure интерфейс для взаимодействия с фигурами
type Figure interface {
	GetSides() []float64
	Visit(visitor) interface{}
}
