package main

import "fmt"

type formatter1 interface {
    format() string
}

type plainText struct {
    message string
}

type bold struct {
    message string
}

type code struct {
    message string
}

func(pt plainText) format() string {
    return pt.message
}

func(b bold) format() string {
    return "**"+b.message+"**"
}

func(c code) format() string {
    return "`"+c.message+"`"
}

func sendMessage(formatter formatter1) string {
    return formatter.format()
}

func main() {
    pt := plainText{message: "ECcentricOG"}
    b := bold{message: "ECcentricOG"}
    c := code{message: "ECcentricOG"}
    fmt.Println(sendMessage(pt))
    fmt.Println(sendMessage(b))
    fmt.Println(sendMessage(c))
}
