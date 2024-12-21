package main

import "fmt"

type authenticationInfo struct {
    username string
    password string
}

func(a authenticationInfo) getBasicAuth() string{
    return "Authentication: Basic "+a.username+":"+a.password
}

func main() {
    user := authenticationInfo{
        username: "ECcentricOG",
        password: "12345678",
    }
    fmt.Println(user.getBasicAuth())
}
