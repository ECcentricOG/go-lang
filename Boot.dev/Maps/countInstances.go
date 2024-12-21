package main

import "fmt"

func getCounts(messagedUsers []string, validUsers map[string]int) {
    for _, user := range messagedUsers {
        if _, ok := validUsers[user]; ok {
            validUsers[user] = validUsers[user] + 1
        }
    }
    fmt.Println(validUsers)
}


func main() {
    mp := map[string]int {
        "cerci" : 0,
        "jaimi" : 0,
        "umesh" : 0,
    }
    sli := []string{"cerci", "umesh","umesh", "umesh", "umesh", "jaimi", "cerci"}
    getCounts(sli, mp)
}
