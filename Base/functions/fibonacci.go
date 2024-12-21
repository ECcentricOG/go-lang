package main

import "fmt"

func main() {
	
	fmt.Println(fibonacci(9))
}

func fibonacci(n int) int {

	if n == 0 {
		return 0
	}	

	if n == 1 {
		return 1
	}

	ans := fibonacci(n - 1) + fibonacci(n - 2)
	return ans
}
