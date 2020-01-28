package crane

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
)

// Crane crane interface
type Crane interface {
	FillUp(volume int) (waterVolume int, err error)
}

type crane struct {
	bath *models.Bath
}

func (c *crane) FillUp(addingVolume int) (waterVolume int, err error) {
	if (c.bath.Volume + addingVolume) > c.bath.MaxVolume {
		err = errors.New("can't add this amount of water. Will be higher than max")
		waterVolume = c.bath.Volume
		return
	}
	c.bath.Volume += addingVolume
	waterVolume = c.bath.Volume
	return
}

func NewCrane(bath *models.Bath) Crane {
	return &crane{
		bath: bath,
	}
}
