package main 

import "fmt"

func main() {

	fmt.Println(greatestNumber(5,1,4,53,6,75,3,5,6))	
}

func greatestNumber(values ...uint) uint {

	greatest := values[0]

	for _,val := range values {
		if val > greatest{
			greatest = val
		}
	}
	return greatest
}
