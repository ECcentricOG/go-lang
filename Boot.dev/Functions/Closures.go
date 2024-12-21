// Function can return a function
package main

import "fmt"

func main() {
    fmt.Println(getEvenNumber())
    fmt.Println(getEvenNumber())
    fmt.Println(getEvenNumber())
}

func getEvenNumber() func() int { // here getEvenNumber function returns another function
    count := 0
    return func() int {
        return count + 2
    }
    
}
