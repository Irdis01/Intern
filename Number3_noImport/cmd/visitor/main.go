package main

import (
	"fmt"

	"github.com/Irdis01/Intern/Number3_noImport/pkg/figure"
	"github.com/Irdis01/Intern/Number3_noImport/pkg/visitor/area"
)

func main() {
	rectangle := figure.NewRectangle(1, 2)
	triangle := figure.NewTriangle(3, 4, 5)
	areaVisitor := area.NewAreaCalculator(triangle)
	fmt.Println(triangle.Visit(areaVisitor).(float64))
	areaVisitor = area.NewAreaCalculator(rectangle)
	fmt.Println(rectangle.Visit(areaVisitor).(float64))
}
