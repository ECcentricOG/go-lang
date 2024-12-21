package main

import (
	"fmt"
	"math"
)

type reactangle struct {
    length, breadth int
}

type circle struct {
    radius float64
}

func(r reactangle) area() int {
    return r.length * r.breadth
}

func(c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func main() {
    rect := reactangle{
        length: 10,
        breadth: 20,
    }
    fmt.Println(rect.area())
    cir := circle{radius: 20,}
    fmt.Println(cir.area())
}
