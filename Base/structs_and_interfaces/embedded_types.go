package main

import "fmt"

func main() {

    per1 := Person{
        Name : "Umesh Unde",
        Address : Address{
            HouseNo : 2017,
            Street : "Bagwan Nagar",
            City : "Shirur",
            PinCode : 412210,
        },
    }

    per2 := Person{
        "ECcentricOG",
        Address{111,"Streeet 2","Pune",414110,},
    }

    fmt.Println(per1)
    fmt.Println(per2)
}
type Person struct {

	Name string
	Address		// Embedded Struct
}

type Address struct {

    HouseNo int
    Street string
    City string
    PinCode int
}

