package main

import "fmt"

func main() {
    fmt.Println(yearsUntilEvents(4))
    fmt.Println(yearsUntilEvents(18))
    fmt.Println(yearsUntilEvents(22))
}

func yearsUntilEvents(age int) (int, int, int) {
    yearsUntilAdult := 18 - age
    if yearsUntilAdult < 0 {
        yearsUntilAdult = 0
    }
    yearsUntilDrinking := 21 - age
    if yearsUntilDrinking < 0 {
        yearsUntilDrinking = 0
    }
    yearsUntilCarRental := 25 - age
    if yearsUntilCarRental < 0 {
        yearsUntilCarRental = 0
    }
    return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}
