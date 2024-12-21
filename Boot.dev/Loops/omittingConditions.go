package main

import "fmt"

func maxMessages(thresh int) int {
    totalCost := 0
    for i:=0; i < thresh; i++ { // if no condition (i < thresh) then infinite loop
        totalCost += 100 + i
    }
    return totalCost
}

func main() {
    fmt.Println(maxMessages(103))
}
