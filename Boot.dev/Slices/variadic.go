package main

import "fmt"

func sum(nums ...int) int {     // variadic args
    sum := 0
    for i:=0; i<len(nums); i++ {
        sum += nums[i]
    }
    return sum
}

func main() {
    arr := []int{1, 2, 3}
    fmt.Println(sum(arr...))    // this is spread operator
    fmt.Println(sum(1,2,3,4,5))
    fmt.Println(sum(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15))
}
