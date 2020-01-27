package figure

type visitor interface {
	VisitTriangle() interface{}
	VisitRectangle() interface{}
}

// Figure интерфейс для взаимодействия с фигурами
type Figure interface {
	GetSides() []float64
	Visit(visitor) interface{}
}
