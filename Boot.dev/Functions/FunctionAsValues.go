package main

import "fmt"

func main() {
    fmt.Println(reformatter("Hello", addExclamation))
    fmt.Println(reformatter("Hello", addPeriod))
}

func reformatter(str string, formatter func(string) string) string {
    fomattedString := formatter(str)
    return "TEXTIO: " + fomattedString
}

func addPeriod(s string) string {
    return s + "."
}

func addExclamation(s string) string {
    return s + "!"
}
