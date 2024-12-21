package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiry string) ([3]string, [3]int) {
    msg := [3]string{primary, secondary, tertiry}
    var arr [3]int
        arr[0] = len(msg[0])
    for i:=1; i<len(msg); i++ {
        arr[i] = arr[i-1] + len(msg[i])
    }
    return msg, arr
}

func main() {
    var arr [3]int
    fmt.Println(arr)    // [0,0,0]
    arr2 := [3]int{1, 2, 3}
    fmt.Println(arr2)    // [1,2,3]
    
    fmt.Println(getMessageWithRetries("Hello sir/madam can I interest you in a yacht?","Please I'll even give you an Amazon gift card?","You're missing out big time"))
    fmt.Println(getMessageWithRetries("Put that coffee down!","Coffee is for closers","Always be closing"))
}
