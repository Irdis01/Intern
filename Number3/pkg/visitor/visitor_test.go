package visitor

import (
	"testing"

	"github.com/Irdis01/Intern/Number3/pkg/figure"
	"github.com/stretchr/testify/assert"
)

const (
	visitortTest = "visitor test"
)

func TestVisitor(t *testing.T) {
	cases := []float64{6, 2}
	testFigure := []figure.Figure{figure.NewTriangle(3, 4, 5), figure.NewRectangle(1, 2)}
	visitor := NewVisitor()
	t.Run(visitortTest, func(t *testing.T) {
		for i := 0; i < len(cases); i++ {
			assert.Equal(t, cases[i], testFigure[i].Visit(visitor))
		}
	})

}
