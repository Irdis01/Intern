package crane

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FillUpTest = "Full up test"
)

func TestCrane_FillUp(t *testing.T) {
	baths := []*models.Bath{{
		MaxVolume: 10,
		Volume:    0,
	},
		{
			MaxVolume: 9,
			Volume:    5,
		},
	}
	expect := []struct {
		result int
		err    error
	}{
		{
			result: 5,
			err:    nil,
		},
		{
			result: 5,
			err:    errors.New("can't add this amount of water. Will be higher than max"),
		},
	}
	t.Run(FillUpTest, func(t *testing.T) {
		for i, _ := range expect {
			testChandelier := NewCrane(baths[i])
			res, err := testChandelier.FillUp(5)
			assert.Equal(t, expect[i].result, res)
			assert.Equal(t, expect[i].err, err)
		}
	})
}
