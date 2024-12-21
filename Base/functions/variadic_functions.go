package main

import "fmt" 

func main() {

	fmt.Println(add(1,5,2,3,4,6,7,8,9,10))
}

func add(args ...int) int {
	
	total := 0
	for _,val := range args {
		total += val
	}
	return total
}
