package main

import "fmt"

type Membership struct {
    Type string
    MessageCharLimit int
}

type User struct {
    Name string
    Membership
}

func newUser(name string, membershipType string) User {
    if membershipType == "premium" {
        return User{Name: name, Membership: Membership{Type: membershipType, MessageCharLimit: 1000,},}
    }
    return User{Name: name, Membership: Membership{Type: membershipType, MessageCharLimit: 100,},}
}

func main() {
    fmt.Println(newUser("Sly", "standard"))
    fmt.Println(newUser("Pattern", "premium"))
    fmt.Println(newUser("Pattern", "standard"))
}
