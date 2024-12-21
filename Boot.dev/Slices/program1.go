package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
    var message []Message
    for i:=0; i<len(messages); i++ {
       // switch msg := messages[i].(type) {
       // case TextMessage : fmt.Println("Text Message")
       // case LinkMessage: fmt.Println("Link Message")
       // case MediaMessage: fmt.Println("Media Message")
       // default : fmt.Println("Other", msg)
       // }
        if filterType ==  messages[i].Type() {
            message = append(message, messages[i])
        }
    }
    return message
}

func main() {
    messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}
    fmt.Println(len(filterMessages(messages, "text")))
    fmt.Println(len(filterMessages(messages, "media")))
    fmt.Println(len(filterMessages(messages, "link")))
    fmt.Println(len(filterMessages(messages, "other")))
}
