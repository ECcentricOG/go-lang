package main

import "fmt"

func main() {

	sli := []int{1,2,3,4,5,6,7,8,9,10}

	add(sli)
}
func add(sli []int) {

	total := 0
	for _,val := range sli {
		total += val
	}
	fmt.Println(total)
}
