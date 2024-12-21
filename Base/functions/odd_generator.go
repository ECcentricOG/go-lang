package main

import "fmt"

func main() {

	nextOddValue := oddGenerator()

	fmt.Println(nextOddValue())
	fmt.Println(nextOddValue())
	fmt.Println(nextOddValue())
	fmt.Println(nextOddValue())
	fmt.Println(nextOddValue())
}

func oddGenerator() func() int {
	
	odd := 1

	return func() (ans int){

		ans = odd
		odd += 2
		return
	}
}
