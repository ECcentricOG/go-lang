package main

import "fmt"

func main() {

	x := 10
	zoo(x)
	fmt.Println(x)		// x = 10
	foo(&x)
	fmt.Println(x)		// x = 0
	
}

func foo(xPtr *int) {
	*xPtr = 0
}

func zoo(x int) {
	x = 0
}
