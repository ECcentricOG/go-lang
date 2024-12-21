package main

import "fmt"

func bultSend(numMessages int) float64 {
    var totalCost float64
    for i:=0; i<numMessages; i++ {
        totalCost += 1 + float64(i)/100.0
    }
    return totalCost
}

func main() {
    fmt.Println(bultSend(10))
    fmt.Println(bultSend(20))
}
