package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
    if len(names) != len(phoneNumbers) {
        return nil, errors.New("Invalid  sizes")
    }
    userMap := make(map[string]user)
    for i := range names {
        userMap[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]} 
    }
    return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
    names := []string{"Eren", "Armin", "Mikasa"}
    phoneNumbers := []int{14355550987, 98765550987, 18265554567}
    fmt.Println(getUserMap(names, phoneNumbers))

    names = []string{"Eren", "Armin"}
    fmt.Println(getUserMap(names, phoneNumbers))
}
