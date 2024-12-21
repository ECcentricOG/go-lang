// if a struct implements all methods from interface then it impliments that interface

package main

import (
	"fmt"
	"math"
)

type Shape interface {

	//Methods
	area() float64
	perimeter() float64
}

type Circle struct {
	Radius float64
}

type Reactangle struct {
	Length, Breadth float64
}

func (r Reactangle) area() float64 {

	return r.Length * r.Breadth
}

func (r Reactangle) perimeter() float64 {

	return 2 * (r.Length + r.Breadth)
}

func (c Circle) area() float64 {

	return math.Pi * c.Radius * c.Radius
}

func (c Circle) perimeter() float64 {

	return 2 * math.Pi * c.Radius
}

func main() {

	var shape1 Shape
	var shape2 Shape

	shape1 = Circle{5}
	shape2 = Reactangle{10, 20}

	fmt.Println(shape1.area())
	fmt.Println(shape1.perimeter())

	fmt.Println(shape2.area())
	fmt.Println(shape2.perimeter())
}
