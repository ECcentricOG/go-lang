package main

import "fmt"


type user struct {
    name string
    number int
}

type sender struct {
    user
    rateLimit int
}

func main() {
    user1 := sender{
        rateLimit: 100,
        user: user{
            name : "Umesh",
            number: 776733,
        },
    }

    fmt.Println(user1.name)     // Embedded Structs can be used to top level
    fmt.Println(user1.number)   // like sender.name if it's nested then is will be like
    fmt.Println(user1.rateLimit)// sender.user.name
}
