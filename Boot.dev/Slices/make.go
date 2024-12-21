package main

import "fmt"

func getMessageCost(messages []string) []float64 {
    messagesCost := make([]float64, len(messages))
    // cap(messagesCost) returns capacity of messages -> len(messages)
    for i:=0; i<len(messages); i++ {
        messagesCost[i] = float64(len(messages[i])) * .01
    }
    return messagesCost
}

func main() {
    fmt.Println(getMessageCost([]string{"Welcome to the movies!", "Enjoy your popcorn!"}))
    fmt.Println(getMessageCost([]string{"I don't want to be here anymore", "Can we go home?", "I'm hungry", "I'm bored"}))
}
