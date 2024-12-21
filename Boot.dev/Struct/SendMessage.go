package main

import "fmt"

func(u User) sendMessage(message string, messageLength int)(bool, string) {
    if messageLength > u.MessageCharLimit {
        return false, ""
    }
    return true, message
}

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
    user1 := User{Name: "Umesh Unde", Membership: Membership{Type: "premium", MessageCharLimit: 1000},}
    fmt.Println(user1.sendMessage("Hello, How are you?", 18))
}

