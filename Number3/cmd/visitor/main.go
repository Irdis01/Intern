package main

import (
	"fmt"

	"github.com/Irdis01/Intern/Number3/pkg/visitor/area"
	"github.com/Irdis01/Intern/Number3/pkg/visitor/figure"
)

func main() {
	rectangle := figure.NewRectangle(1, 2)
	triangle := figure.NewTriangle(3, 4, 5)
	areaVisitor := area.NewAreaCalculator()
	fmt.Println(triangle.Visit(areaVisitor).(float64))
	fmt.Println(rectangle.Visit(areaVisitor).(float64))
}
