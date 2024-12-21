package main

import "fmt"

func main() {

	a := 20
	fmt.Println(&a)		// Memory address of a
	x := new(int)
	fmt.Println(x)
	foo(x)
	fmt.Println(x)
}

func foo(xPtr *int) {
	*xPtr = 10
}
