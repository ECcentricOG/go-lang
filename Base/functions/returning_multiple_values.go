package main

import "fmt"

func main() {

	fmt.Print("Enter Dividend : ")
	var dividend int 
	fmt.Scanln(&dividend)
	fmt.Print("Enter Divisor")
	var divisor int
	fmt.Scanln(&divisor)

	ans , remainder := divide(dividend,divisor)
	fmt.Println("Ans : ",ans,"\n Remainder : ",remainder)
}

func divide(dividend int ,divisor int) (int, int) {
	ans := dividend / divisor
	remainder := dividend % divisor
	return ans, remainder
}
