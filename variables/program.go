package main

import "fmt"

func main() {
	var (
		a = 5
		b = 20
	)
	fmt.Println(a + b)
	const pi float32 = 3.1415
	var str string
	str = "ECcentric"
	fmt.Println(str)
	str = "OG"
	fmt.Println(str)
	var result string = "ECcentric"	 + str
	fmt.Println(result)
	fmt.Println("ECcentricOG" == result)
	fmt.Println("Value of Pi is ", pi)
	fmt.Println("Hello I am ", result)
}
