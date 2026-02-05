package main

import "fmt"

func main() {
	var principal float64
	var rate float64
	var time float64

	fmt.Println("Enter Principal amount:")
	fmt.Scan(&principal)
	fmt.Println("Enter Rate of Interest:")
	fmt.Scan(&rate)
	fmt.Println("Enter Time in years:")
	fmt.Scan(&time)

	si := (principal * rate * time) / 100

	fmt.Printf("Simple Interests is : %f\n", si)
}
