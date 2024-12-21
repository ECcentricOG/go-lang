package main

import "fmt"

func main() {
    fmt.Println(getProductMessage("basic"))
    fmt.Println(getProductMessage("premium"))
    fmt.Println(getProductMessage("enterprise"))
    fmt.Println(getProductMessage("free"))
}

func getProductMessage(tier string) string {
    quantity, price, _ := getProductInfo(tier)      // here '_' is used to ignore the return value cause go compiler throws error if variable is not used.
    return quantity + price
}

// this func return 3 values of string datatypes
func getProductInfo(tier string) (string, string, string) {
    if tier == "basic" {
        return "1000 text per month", "$30 per month",  "most popular"
    } else if tier == "premium" {
        return "50000 text per month", "$60 per month", "best value"
    } else if tier == "enterprise" {
        return "unlimited text per month", "$100 per month", "customizable"
    } else {
        return " ", " ", " "
    }
}
