package main 

import "fmt"

func main() {
	fmt.Print("Enetr a no : ")
	var number int
	fmt.Scanln(&number)

	if number%3 == 0 && number%5 == 0{
		fmt.Println("FizzBuzz")
	}else if number%3 == 0 {
		fmt.Println("Fizz")
	}else if number%5 == 0 {
		fmt.Println("Buzz")
	}
}
