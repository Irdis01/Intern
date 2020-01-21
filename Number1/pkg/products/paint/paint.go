package paint

type ProductPaint interface {
	GetName() (name string)
	GetType() (name string)
}

type paint struct {
	name string
}

func (p *paint) GetName() (name string) {
	name = p.name
	return name
}

func NewPaint(name string) *paint {
	newPaint := paint{name: name}
	return &newPaint
}

func (p *paint) GetType() (name string) {
	name = "paint"
	return
}
