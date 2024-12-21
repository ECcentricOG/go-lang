package main

import (
	"fmt"
	"unicode"
)

func isValidPassword(password string) bool {
    if len(password) < 5 || len(password) > 12 {
        return false
    }

    hasUpperCase := false
    hasDigit := false
    for _, char := range password {
        if unicode.IsUpper(char) {
            hasUpperCase = true 
        }

        if unicode.IsDigit(char) {
            hasDigit= true
        }

        if hasDigit && hasUpperCase {
            break
        }
    }
    return hasDigit && hasUpperCase
}

func main() {
    fmt.Println(isValidPassword("password123"))
    fmt.Println(isValidPassword("Password123"))
    fmt.Println(isValidPassword("pass12"))
    fmt.Println(isValidPassword("Pass12"))
    fmt.Println(isValidPassword("ECcentricOG3"))
}
