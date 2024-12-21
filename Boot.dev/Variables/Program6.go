package main

import (
    "fmt"
)

func main() {
    var username string = "ECcentricOG"
    var password string = "12345678"
    msg := username+":"+password        // '+' is a concat operator
    fmt.Println(msg)
}
