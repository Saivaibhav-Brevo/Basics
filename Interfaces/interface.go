package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Square struct {
	side float64
}

func (r *Rectangle) Area(w, h float64) (float64, error) {
	return (w * h), nil
}

func (r *Square) Area(s float64) (float64, error) {
	return math.Pow(s, 2), nil
}

func main() {

	rect := Rectangle{}

	res, _ := rect.Area(10.5, 11.5)

	fmt.Println("Area of Rectangle is: ", res)

}
