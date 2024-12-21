package main

import "fmt"

func monthlyBillIncrease(costPerSecond, numLastMonth, numThisMonth int) int {
    var lastMonthBill int
    var thisMonthBill int
    lastMonthBill = getBillForMonth(costPerSecond, numLastMonth)
    thisMonthBill = getBillForMonth(costPerSecond, numThisMonth)
    return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
    return costPerSend * messagesSent
}

func main() {
    fmt.Println(monthlyBillIncrease(2, 89, 102))
    fmt.Println(monthlyBillIncrease(2, 98, 104))
    fmt.Println(monthlyBillIncrease(3, 50, 40))
    fmt.Println(monthlyBillIncrease(3, 60, 60))
}
