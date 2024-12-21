package main

import "fmt"

type employee interface {
    getName() string
    getSalary() int
}

type contractor struct {
    name string
    hourlyPay int
    hoursPerYear int
}

func(c contractor) getName() string {
    return c.name
}

func(c contractor) getSalary() int {
    return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
    name string
    salary int
}

func(ft fullTime) getName() string {
    return ft.name
}

func(ft fullTime) getSalary() int{
    return ft.salary
}

func displayInfo(e employee) {
    fmt.Println("Salary : ",e.getSalary())
    fmt.Println("Name: ",e.getName())
}

func main() {
    cont := contractor{name: "Useless", hourlyPay: 250, hoursPerYear: 200}
    fullT := fullTime{name: "Umesh", salary: 30000} 
    displayInfo(cont)
    displayInfo(fullT)
}
