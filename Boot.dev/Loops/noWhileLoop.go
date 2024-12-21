package main

import "fmt"

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
    actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance >= 0{
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func main() {
    fmt.Println(getMaxMessagesToSend(1.1, 5))
    fmt.Println(getMaxMessagesToSend(1.3, 10))
}