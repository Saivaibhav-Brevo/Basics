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

func main() {
	emp1 := employee{
		name:   "Vaibhav",
		age:    22,
		salary: 50000.0,
	}
	fmt.Println("employee details:", emp1)

	showDetails(&emp1)
	emp1.output()
}

func showDetails(emp *employee) {
	fmt.Println("employee name:", emp.name)
	fmt.Println("employee age:", emp.age)
	fmt.Println("employee salary:", emp.salary)
}
