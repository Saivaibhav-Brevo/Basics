package main

import "fmt"

type Person struct {
	id int
}

func (p *Person) assignement() error {
	fmt.Println("Function calleed")
	return nil
}

func main() {
	p1 := Person{}
	err := p1.assignement()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
