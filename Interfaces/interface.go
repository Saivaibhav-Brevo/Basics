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

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Square) Area() float64 {
	return math.Pow(r.side, 2)
}

func main() {

	rect := Rectangle{width: 10.5, height: 11.5}

	fmt.Println("Area of Rectangle is: ", rect.Area())

}
