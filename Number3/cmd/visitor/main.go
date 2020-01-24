package main

import (
	"fmt"

	myVisitor "github.com/Irdis01/Intern/Number3/pkg/visitor/area"
	"github.com/Irdis01/Intern/Number3/pkg/visitor/figure"
)

func main() {
	rectangle := figure.NewRectangle(1, 2)
	triangle := figure.NewTriangle(3, 4, 5)
	visitor := myVisitor.NewAreaCalculator()
	fmt.Println(triangle.Visit(visitor).(float64))
	fmt.Println(rectangle.Visit(visitor).(float64))
}
