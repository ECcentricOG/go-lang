package main

import (
	"fmt"
	"time"
)

type message interface {
    getMessage() string
}

type birthdayMessage struct {
    birthdayTime time.Time
    recipientName string
}

func(bm birthdayMessage) getMessage() string {
    return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
    numberOfSends int
    reportName string
}

func(sr sendingReport) getMessage() string {
    return fmt.Sprintf("Your %s report is ready. You've sent %d messages.", sr.reportName, sr.numberOfSends)
}

func sendMessage(msg message) (string, int) {
    return msg.getMessage(), len(msg.getMessage()) * 3
}

func main() {
    fmt.Println(sendMessage(sendingReport{reportName: "First Report", numberOfSends: 123,}))
    fmt.Println(sendMessage(birthdayMessage{birthdayTime: time.Now(), recipientName: "Umesh Unde"}))
}
