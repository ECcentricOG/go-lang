package main

import "fmt"

func main() {
    arr := [10]int{1,2,3,4,5,6,7,8,9,0}
    sli := arr[2:8]
    fmt.Println("Length of Slice : ",len(sli))
    fmt.Println("Length of Array : ",len(arr))
    fmt.Println("Capacity of Array : ",cap(arr))
    fmt.Println("Capacity of Slice : ",cap(sli))    // capacity is calculated by underlying array of slice by arr_length - slice_start_ele
}
