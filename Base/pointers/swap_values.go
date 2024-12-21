package main

import "fmt"

func main() {

	x := 10
	y := 20
	fmt.Println("x and y are",x,y)
	fmt.Println("Memory Address of x and y : ",&x,&y)
	swap(&x,&y)
	fmt.Println("x and y are",x,y)
	fmt.Println("Memory Address of x and y : ",&x,&y)
}

func swap(xPtr ,yPtr *int) {

	temp := *xPtr
	*xPtr = *yPtr
	*yPtr = temp
}
