package main

import (
	"errors"
	"fmt"
)

type Person struct {
	id int
}

func (p *Person) assignement() error {
	fmt.Println("Reciever type Function calleed")
	return errors.New("This is an error")
}

func compute(p *Person) error {
	errorFunc := p.assignement()
	return errorFunc
}

func main() {
	p1 := Person{}

	err := compute(&p1)

	if err != nil {
		fmt.Println("Error:", err)
	}

}
