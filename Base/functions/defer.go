package main

import "fmt"

func main() {
	defer one()
	two()
	three()			// one() will execute at the end after two() and three()
				// if all are differ the first defer will exe last and second will exe second last
}

func one() {
	fmt.Println("In function one")
}

func two() {
	fmt.Println("In function two")
}

func three() {
	fmt.Println("In function three")
}
