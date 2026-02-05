package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	var p1 person
	p1 = person{
		firstName: "Vaibhav",
		age:       12,
		lastName:  "Kumar",
	}

	fmt.Println(p1.firstName)

}
