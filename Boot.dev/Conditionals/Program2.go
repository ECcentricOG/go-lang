package main

import "fmt"

func main() {
    password := "12345678"
    if length := len(password); length < 8 {
        fmt.Println("Short password")
    } else if length > 50 {
        fmt.Println("Large password")
    } else {
        fmt.Println("Good to go")
    }
}
