package main

import "fmt"

func main() {
	fmt.Print("Enter a no : ")
	var number uint
	fmt.Scanln(&number)
	
	fmt.Println(factorial(number))
}

func factorial(number uint) uint {
	if number == 0 {
		return 1
	}

	return number * factorial(number - 1)
}
