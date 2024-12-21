package main

import "fmt"

func main() {
    const name = "Umesh Unde"
    const openRate = 30.54
    msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate)
    fmt.Println(msg)
}
