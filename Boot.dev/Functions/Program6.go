package main

import (
	"errors"
	"fmt"
)

func main() {
    fmt.Println(calculator(10, 20))
    fmt.Println(calculator(30, 0))
    fmt.Println(calculator(40, 10))
}

func calculator(a, b int) (mul, div int, err error) {   //named return values these (mul, div) are variable so no need to declare again
    if b == 0 {
        return 0, 0, errors.New("Can't divide by zero")
    }
    mul = a * b
    div = a / b
    return mul, div, nil    // nil is a zero value for an error like null
}














