package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary float64
}

// reciever type
func (emp *employee) output() {
	fmt.Println("employee name:", emp.name)
	fmt.Println("employee age:", emp.age)
	fmt.Println("employee salary:", emp.salary)
}

func (emp *employee) input() {
	fmt.Println("Enter employee details: ")
	fmt.Print("Name: ")
	fmt.Scanln(&emp.name)

	fmt.Print("Age: ")
	fmt.Scanln(&emp.age)

	fmt.Print("Salary: ")
	fmt.Scanln(&emp.salary)
}

func main() {
	// emp1 := employee{
	// 	name:   "Vaibhav",
	// 	age:    22,
	// 	salary: 50000.0,
	// }

	var emp1 employee

	emp1.input()
	fmt.Println("employee details:", emp1)

	// showDetails(&emp1)
	emp1.output()
}

func showDetails(emp *employee) {
	fmt.Println("employee name:", emp.name)
	fmt.Println("employee age:", emp.age)
	fmt.Println("employee salary:", emp.salary)
}
