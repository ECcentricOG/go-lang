package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("no dividing by zero")
    }
    return x/y, nil
}

func main() {
    fmt.Println(divide(10,2))
    fmt.Println(divide(10,0))
    fmt.Println(divide(15,5))
    fmt.Println(divide(100,0))
}
