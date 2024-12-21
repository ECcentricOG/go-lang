package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
    for i := range messages {
        messages[i].tags = tagger(messages[i])
    }
    return messages
}

func tagger(msg sms) []string {
	tags := []string{}

    if strings.Contains(strings.ToLower(msg.content), "sale") {
        tags = append(tags, "Promo")
    }

    if strings.Contains(strings.ToLower(msg.content), "urgent") {
        tags = append(tags, "Urgent")
    }
    return tags
}

func main() {
    messages:= []sms{{id: "001", content: "Urgent, please respond!"}, {id: "002", content: "Big sale on all items!"}}
    fmt.Println(tagMessages(messages, tagger))
}
