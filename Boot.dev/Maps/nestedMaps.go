package main

import (
	"fmt"
)

func getNameCounts(names []string) map[rune]map[string]int {
    usersMap := make(map[rune]map[string]int)
    for _, name := range names {
        runes := []rune(name) 
        if _, ok := usersMap[runes[0]]; ok {
            if _, ok := usersMap[runes[0]][name]; ok {
                usersMap[runes[0]][name] =  usersMap[runes[0]][name] + 1 
            } else {
                usersMap[runes[0]][name] = 1
            }
        } else {
            entry := map[string]int{ name : 1}
            usersMap[runes[0]] = entry
        }
    }
    return usersMap
}

func main() {
    fmt.Println(getNameCounts([]string{"umesh", "unde", "umesh", "umesh", "sakshi", "ppoja", "pooja", "shruti", "unde", "shritika"}))
}
