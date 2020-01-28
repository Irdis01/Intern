package bath

type Bath interface {
	FillUp(volume int)
}

type bath struct {
	waterVolume int
}

func (b *bath) FillUp(volume int) {
	b.waterVolume = volume
}

func NewBath() Bath {
	return &bath{}
}
