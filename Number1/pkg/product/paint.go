package product

type paint struct {
	name string
}

func (p *paint) GetName() (name string) {
	name = p.name
	return name
}

//NewPaint фабрика красок
func NewPaint(name string) Product {
	return &paint{
		name: name,
	}
}

func (p *paint) GetType() (name string) {
	name = "paint"
	return
}
