package main

import "fmt"

func main() {

	react := Reactangle{10,20}
	fmt.Println(react.area())
}

type Reactangle struct {
	length , breadth float64
}

// only available fot Reactangle obj
func (r *Reactangle) area() float64 {
	return r.length * r.breadth
}
