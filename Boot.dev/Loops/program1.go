package main

import "fmt"

func countConnections(groupSize int) int {
    return ((groupSize*(groupSize+1)) / 2) - groupSize
}

func main() {
    fmt.Println(countConnections(1))
    fmt.Println(countConnections(2))
    fmt.Println(countConnections(3))
    fmt.Println(countConnections(4))
    fmt.Println(countConnections(0))
    fmt.Println(countConnections(10))
    fmt.Println(countConnections(100))
    fmt.Println(countConnections(1000))
}
