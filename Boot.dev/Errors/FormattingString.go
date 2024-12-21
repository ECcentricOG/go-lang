package main

import "fmt"

func getSMSErrorString(cost float64, recipitent string) string {
    return fmt.Sprintf("SMS that costs $%.2f to be sent to '%s' can not be sent",cost, recipitent)
}

func main() {
    fmt.Println(getSMSErrorString(1.4, "+1 (435) 555 0923"))
    fmt.Println(getSMSErrorString(2.1, "+2 (702) 555 3452"))
    fmt.Println(getSMSErrorString(32.1, "+1 (801) 555 7456"))
    fmt.Println(getSMSErrorString(14.4, "+1 (234) 555 6545"))
}
