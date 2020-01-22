package paint

//ProductPaint интерфейс карсок
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

//NewPaint фабрика красок
func NewPaint(name string) ProductPaint {
	return &paint{
		name: name,
	}
}

func (p *paint) GetType() (name string) {
	name = "paint"
	return
}
