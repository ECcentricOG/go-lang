// General rule do not use panic it prints stacktrace like java exceptions
package main

import (
	"errors"
	"fmt"
	//"log"
)

// panic and recover() are like try catch

type User struct {
    name string
    id int
}

func getUser(id int) (User, error){
    if id < 10{
        return User{id: id}, nil
    }
    return User{}, errors.New("No Such User")
}

func enrichUser(userId int) User {
    user, err := getUser(userId)
    if err != nil {
        //log.Fatal("Log : User not found")
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recover From Panic : ", r)
        }
    }() // this is anonymous function decleared last '()' is for so we can call it
    fmt.Println(enrichUser(1))
    fmt.Println(enrichUser(12))
    fmt.Println(enrichUser(3))
}
