package main

import "fmt"

type messageToSend struct {
    message string
    sender user
    recipient user
}

type user struct {
    name string
    number int
}

func canSendMessage(mToSend messageToSend) bool {
    if mToSend.sender.name == "" || mToSend.sender.number == 0 {
        return false
    }
    if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
        return false
    }
    return true
}

func main() {
    fmt.Println(canSendMessage( messageToSend {
        message: "Hello GOAT",
        sender : user{name : "Umesh Unde", number : 3},
        recipient : user{name : "ECcentricOG", number : 33},
    }))


    fmt.Println(canSendMessage( messageToSend {
        message: "Hello GOAT",
        sender : user{name : "Umesh Unde"},
        recipient : user{name : "ECcentricOG", number : 33},
    }))


    fmt.Println(canSendMessage( messageToSend {
        message: "Hello GOAT",
        sender : user{name : "Umesh Unde", number : 3},
        recipient : user{ number : 33},
    }))
}
