package main

import "fmt"

func sendSMS(message string) (int, error) {
    const maxTextLen = 25
    const costPerChar = 2
    if len(message) > maxTextLen {
        return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
    }
    return costPerChar * len(message), nil
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
    custCost, err := sendSMS(msgToCustomer)
    if err != nil {
        return 0, err
    }
    spouseCost, err := sendSMS(msgToSpouse)
    if err != nil {
        return 0, err
    }
    return custCost + spouseCost, nil
}

func main() {
    fmt.Println(sendSMSToCouple("Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift."))
    fmt.Println(sendSMSToCouple("Thanks for joining us!", "Have a good day."))
}
