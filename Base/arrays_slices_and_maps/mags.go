package main

import "fmt"

func main() {
	// var m1 map[string]int
	// m1["id"] = 1		assignment to entry in nil map
	// maps have to be initialized before they can be used

	m1 := make(map[string]string)
	m1["key"] = "value"
	fmt.Println(m1)
	delete(m1,"key")
	fmt.Println(m1)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Berryllim"
	elements["B"] = "Boron"
	elements["C"] = "Cardon"

	fmt.Println(elements)
	fmt.Println(elements["Li"])

	if element,status := elements["U"] ; status {
		fmt.Println(element,status)
	}else {
		fmt.Println("Element doesn't exist")
	}
	
	var eleName string
	fmt.Print("Enter element name : ")
	fmt.Scanln(&eleName)

	ele := map[string]map[string]string {
		"H" : map[string]string {
			"name" : "Hydrogen",
			"state" : "Gas",
		},
		"He" : map[string]string {
			"name" : "Helium",
			"state" : "Gas",
		},
		"Li" : map[string]string {
			"name" : "Lithium",
			"state" : "Solid",
		},
		"Be" : map[string]string {
			"name" : "Berrylium",
			"state" : "solid",
		},
		"B" : map[string]string {
			"name" : "Boron",
			"state" : "solid",
		},
		"C" : map[string]string {
			"name" : "Carbon",
			"state" : "Solid",
		},
	}
	if elem , sta := ele[eleName] ; sta {
		fmt.Println(elem)
	}else {
		fmt.Println("Element not found")
	}
}
