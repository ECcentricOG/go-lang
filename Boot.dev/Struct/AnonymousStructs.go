package main

import "fmt"

type Car struct {
    brand string
    model string
    mileage float64
    wheel struct {      // wheel is anonymous struct
        radius int
        material string
    }
}


func main() {

    person := struct {  // no name but must be in func body can't be global
        name string
        age int
    } {
        name : "Umesh",
        age : 21,
    }

    wagonR := Car{
        brand: "Suzuki",
        mileage: 33.3,
        wheel: struct{
            radius int
            material string
        }{
            radius: 10,
            material: "Letex",
        },
    }

    fmt.Println(wagonR)
    fmt.Println(person)
}

