package visitor

import (
	"testing"

	"github.com/Irdis01/Intern/Number3/pkg/visitor/figure"
	"github.com/stretchr/testify/assert"
)

func TestVisitor(t *testing.T) {
	cases := []float64{6, 2}
	testFigure := []figure.Figure{figure.NewTriangle(3, 4, 5), figure.NewRectangle(1, 2)}
	visitor := NewAreaCalculator()
	t.Run("visitor test", func(t *testing.T) {
		for i := 0; i < len(cases); i++ {
			assert.Equal(t, cases[i], testFigure[i].Visit(visitor))
		}
	})

}
