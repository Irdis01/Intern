package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	newCastleTest = "newCastleTest"
)

func TestBuilder_NewCastle(t *testing.T) {
	var defenceMap = map[string]int{
		"stone":      10,
		"heavy door": 5,
	}
	expect := 15
	builder := NewBuilder()
	castle := builder.NewCastle()
	t.Run(newCastleTest, func(t *testing.T) {
		assert.Equal(t, castle.CalculateDefence(defenceMap), expect)
	})
}

func TestBuilder_NewHouse(t *testing.T) {
	var defenceMap = map[string]int{
		"wood":       5,
		"light door": 2,
	}
	expect := 7
	builder := NewBuilder()
	castle := builder.NewHouse()
	t.Run(newCastleTest, func(t *testing.T) {
		assert.Equal(t, castle.CalculateDefence(defenceMap), expect)
	})
}

func TestBuilder_NewTent(t *testing.T) {
	var defenceMap = map[string]int{
		"cloth": 2,
		"none":  0,
	}
	expect := 2
	builder := NewBuilder()
	castle := builder.NewTent()
	t.Run(newCastleTest, func(t *testing.T) {
		assert.Equal(t, castle.CalculateDefence(defenceMap), expect)
	})
}
