package main

import "fmt"

func main() {

	fmt.Print("Enter a No : ")
	var tempInFah float64
	fmt.Scanf("%f",&tempInFah)
	
	tempInCel := (tempInFah - 32) * 5/9
	fmt.Println("Temperature in Cel : ",tempInCel)
}
