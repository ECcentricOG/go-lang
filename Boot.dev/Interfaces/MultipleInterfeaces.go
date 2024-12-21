package main

import "fmt"

type expense interface {
    cost() int
}

type formatter interface {
    format() string
}

type email struct {
    isSubscibed bool
    body string
}

func(e email) cost() int {      // Here Implements expense interface
    if e.isSubscibed {
    return 2 * len(e.body)
    }
    return 5 * len(e.body)
}

func(e email) format() string {     // Here Implements formatter interface
    if e.isSubscibed {
        return fmt.Sprintf(`'%s' | Subscribed`, e.body)
    }
    return fmt.Sprintf(`'%s' | Not Subscribed`, e.body)
}

func testExpense(e expense) {
    fmt.Println(e.cost())
}

func testFormatter (e formatter) {
    fmt.Println(e.format())
}

func main() {
    eml1 := email{isSubscibed: true, body: "Beautiful Morning"} 
    eml2 := email{isSubscibed: false, body: "Fein Fein Fein Fein Fein"} 
    testExpense(eml1)
    testExpense(eml2)
    testFormatter(eml1)
    testFormatter(eml2)
}
