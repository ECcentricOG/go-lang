package main

import"fmt"

func getMonthlyPrice(tier string) int {
    if tier == "basic" {
        return 100
    } else if tier == "premium" {
        return 150
    } else if tier == "enterprise" {
        return 500
    } else {
        return 0
    }
}

func main() {
    fmt.Println(getMonthlyPrice("basic"))
    fmt.Println(getMonthlyPrice("premium"))
    fmt.Println(getMonthlyPrice("enterprise"))
    fmt.Println(getMonthlyPrice("free"))
}
