package main

import "fmt"

type divideError struct {
    divident float64 
}

func(de divideError) Error() string {
    return fmt.Sprintf("can not divide %v by zero", de.divident)
}

func divide(dividend, divisor float64) (float64, error) {
    if divisor == 0 {
        return 0, divideError{divident: dividend}
    }
    return dividend / divisor, nil
}

func main() {
    fmt.Println(divide(10, 2))
    fmt.Println(divide(15, 3))
    fmt.Println(divide(10, 0))
    fmt.Println(divide(100, 0))
}