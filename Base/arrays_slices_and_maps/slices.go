package main

import "fmt"

func main() {

	var sli []int
	sli2 := make([]float64,5,10)	// 5 is allocated memory and 10 is total capacity
	fmt.Println(sli,sli2)
	arr := [5] int{1,2,3,4,5}
	x := arr[1:4]			// x is slice of array arr contains [2,3,4]
	fmt.Println(x)
	fmt.Println(len(sli2))	
	slice1 := []int{1,2,3}
	slice2 := append(slice1,4,5)
	fmt.Println(slice1,slice2)

	slice3  := [] int{1,2,3}
	slice4 := make([]int,2)
	
	copy(slice4,slice3)
	fmt.Println(slice3,slice4)
}

