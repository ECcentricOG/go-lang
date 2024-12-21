package main

import "fmt"

func main() {
    bootup()
}

func bootup() {
    defer fmt.Println("Connection Closed")  // it will execute no matter what like a finally block in java
    // defer line execute first but output displays at the end of the function
    fmt.Println("Connection is Open")
    // runtime exception
}
