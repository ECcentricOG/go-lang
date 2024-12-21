package main

import "fmt"

func main() {
	
	a , b := half(2)	
	fmt.Println(a,b)
	c , d := half(7)	
	fmt.Println(c,d)
}

func half(num int) (halve int,isEven bool) {

	halve = num / 2
	
	if num % 2 == 0 {
		isEven = true
	} else {
		isEven = false
	}
	return
}
