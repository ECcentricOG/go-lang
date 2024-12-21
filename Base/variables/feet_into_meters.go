package main

import "fmt"

func main() {
	fmt.Print("Enter a No : ")
	var feet float32
	fmt.Scanf("%f",&feet)

	meter := feet * float32(0.3048)

	fmt.Println("Distance in meter : ",meter)
}
