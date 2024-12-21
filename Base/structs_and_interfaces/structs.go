package main

import "fmt"

func main() {

	emp1 := Employee{name: "Umesh", id: 1, role: "Mobile Developer"}
	var emp2 Employee = Employee{name: "ECcentric", id: 3, role: "Tester"}
	emp3 := new(Employee)
	emp3.name = "Useless"
	emp3.id = 2
	emp3.role = "Looser"
	emp4 := Employee{"UU", 4, "Full Stack"}
	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
	fmt.Println(emp4)
}

type Employee struct {
	name string
	id   int
	role string
}
