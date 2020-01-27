package area

import (
	"testing"

	visitFig "github.com/Irdis01/Intern/Number3_noImport/pkg/figure"
	"github.com/stretchr/testify/assert"
)

const (
	visitortTest = "visitor test"
)

func TestVisitor(t *testing.T) {
	cases := []float64{6, 2}
	testFigure := []visitFig.Figure{visitFig.NewTriangle(3, 4, 5), visitFig.NewRectangle(1, 2)}
	t.Run(visitortTest, func(t *testing.T) {
		for i := 0; i < len(cases); i++ {
			visitor := NewAreaCalculator(testFigure[i])
			assert.Equal(t, cases[i], testFigure[i].Visit(visitor))
		}
	})

}
