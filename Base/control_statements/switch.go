package main

import "fmt"

func main() {
	
	number := 5 

	switch number {
	case 0 : fmt.Println("Zero")
	case 1 : fmt.Println("One")
	case 3 : fmt.Println("Three")
	case 2 : fmt.Println("Two")
	default : fmt.Println("Unknown")
	}
}
