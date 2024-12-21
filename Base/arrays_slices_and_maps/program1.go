package main

import "fmt"

func main() {
	
	x := []int {
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	
	var small int = x[0]

	for _,val := range x {
		if val < small {
			small = val
		}
	}
	fmt.Println(small)
}
