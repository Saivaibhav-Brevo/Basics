package main

import "fmt"

type employee struct {
	name   string
	age    int
	salary float64
}

func main() {
	emp1 := employee{
		name:   "Vaibhav",
		age:    22,
		salary: 50000.0,
	}
	fmt.Println("employee details:", emp1)
}
