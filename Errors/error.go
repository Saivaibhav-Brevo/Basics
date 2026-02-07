package main

import (
	"errors"
	"fmt"
)

type Person struct {
	id int
}

var customError = errors.New("This is a custom error")

func (p *Person) assignement() error {
	fmt.Println("Reciever type Function calleed")
	return customError
}

func compute(p *Person) error {
	errorFunc := p.assignement()
	return errorFunc
}

func main() {
	p1 := Person{}

	err := compute(&p1)

	if err != nil {
		if errors.Is(err, customError) {
			err = fmt.Errorf("failed to load user: %w", err)
			fmt.Println(err)
		}
	}

}
