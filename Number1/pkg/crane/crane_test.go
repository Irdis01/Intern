package crane

import (
	"errors"
	"github.com/Irdis01/Intern/Number1/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	fillUpTestSuccess = "Full up test sucess"
	fillUpTestError   = "Full up test sucess"
	fillUpError       = "can't add this amount of water. Will be higher than max"
)

func TestCrane_FillUpSuccess(t *testing.T) {
	baths := &models.Bath{
		MaxVolume: 10,
		Volume:    0,
	}
	expect := 5
	testCrane := NewCrane(baths)
	t.Run(fillUpTestSuccess, func(t *testing.T) {
		res, err := testCrane.FillUp(5)
		assert.NoError(t, err, "unexpected error")
		assert.Equal(t, expect, res)
	})
}

func TestCrane_FillUpError(t *testing.T) {
	baths := &models.Bath{
		MaxVolume: 9,
		Volume:    5,
	}
	expect := errors.New(fillUpError)
	testCrane := NewCrane(baths)
	t.Run(fillUpTestError, func(t *testing.T) {
		_, err := testCrane.FillUp(5)
		assert.Equal(t, expect, err)
	})
}
