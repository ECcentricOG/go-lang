package main

import "fmt"

func main() {
	nextEven := evenGenerator()

	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

func evenGenerator() func() uint {
	current := uint(0)

	return func() (ret uint) {
		ret = current
		current += 2
		return
	}
}
