package main

import "fmt"

func main() {

	division(24,6)
	division(4,0)
	division(6,3)
	divisionWithRecovrt(4,0)
	divisionWithRecovrt(4,1)
}

func division(num1,num2 int) {

	defer handlePanic()

	if num2 == 0 {
		panic("divide by zero")
	} else {
		fmt.Println(num1/num2)
	}
}

func handlePanic() {

	exp := recover()

	if  exp != nil {
		fmt.Println("Recover : ",exp)
	}
}

func divisionWithRecovrt(num1,num2 int) {
	
	defer func() {
		
		exception := recover()

		if exception != nil {
			fmt.Println(exception)
		}
	}()
	if num2 == 0 {
		panic("divide by zero")
	}else {
		fmt.Println(num1/num2)
	}
}

