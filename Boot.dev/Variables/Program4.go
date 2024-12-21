package main

import (
    "fmt"
)

func main() {
    numMessagesFromUmesh := 72
    costPerMessage := .02
    totalCost := costPerMessage * float64(numMessagesFromUmesh)
    fmt.Printf("Umesh spend %.2f on text messages today\n", totalCost)
}
