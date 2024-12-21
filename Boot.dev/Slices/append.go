package main

import "fmt"

type cost struct {
    value float64
    day int
}

func getCostByDay(costs []cost) []float64 {
    maxDays := []int{}
    for i:=0; i<len(costs); i++ {
        maxDays = append(maxDays, costs[i].day)
    }
    max := 0
    for i:=0; i<len(maxDays); i++ {
        if maxDays[i] > max {
            max = maxDays[i]
        }
    }
    totals := make([]float64, max + 1)
    for i:=0; i<=max; i++ {
        totals[costs[i].day] += costs[i].value 
    }
    return totals
}

func main() {
    costs :=  []cost{
        cost{day :0, value : 1.0},
        cost{day :1, value : 2.0},
        cost{day :5, value : 2.5},
        cost{day :2, value : 3.6},
        cost{day :1, value : 2.7},
        cost{day :1, value : 3.3},
    }
    fmt.Println(getCostByDay(costs))
}
