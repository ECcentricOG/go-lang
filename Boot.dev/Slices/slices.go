// lowIndex is inclusive and highIndex is exclusive.

package main

import (
	"errors"
	"fmt"
)

const (
    planFree = "free"
    planPro = "pro"
    planEnterprise = "enterprise"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
    if plan == planFree {
        return messages[:2], nil
    } else if plan == planPro {
        return messages[:], nil
    } else {
        return nil, errors.New("unsupported plan")
    }
}

func main() {
    arr := [5]int{1,2,3,4,5}    // array
    sli := []int{1,2,3,4,5,6}   // sclice
    fmt.Println(sli)
    fmt.Println(arr[2:4])   // 3 4
    fmt.Println(arr[:4])    // 1 2 3 4
    fmt.Println(arr[2:])    // 3 4 5
    fmt.Println(arr[:])     // 1 2 3 4 5
    
    fmt.Println(getMessageWithRetriesForPlan(planFree, [3]string{"Hello sir/madam can I interest you in a yacht?","Please I'll even give you an Amazon gift card?","You're missing out big time",}))                                
    fmt.Println(getMessageWithRetriesForPlan(planPro, [3]string{"Hello sir/madam can I interest you in a yacht?","Please I'll even give you an Amazon gift card?","You're missing out big time",}))                                
    fmt.Println(getMessageWithRetriesForPlan(planEnterprise, [3]string{"Hello sir/madam can I interest you in a yacht?","Please I'll even give you an Amazon gift card?","You're missing out big time",}))                                
}
