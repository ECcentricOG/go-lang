package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	arr[1] = 20

	fmt.Println(arr)

	var arr2 [5]int = [5]int{1,2,3,4,5}

	fmt.Println(arr2)

	arr3 := [4]int{10,20,30,40}
	
	sum := 0
	for i := 0;i < len(arr3); i++ {
		sum += arr3[i]
	}

	fmt.Println(sum)
}
