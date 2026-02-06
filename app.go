package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var temp [5]int

	var str [2]string

	str[0] = "sai"
	str[1] = "vaibhav"

	fmt.Println("arr is ", str)
	fmt.Println("arr is ", arr)
	fmt.Println("arr is ", temp)

	//slices

	temp2 := []float64{2.4, 5.6, 7.3, 8.9}

	slice1 := temp2[1:3]
	slice2 := temp2[0:4]
	slice3 := temp2[0:]
	slice4 := temp2[:3]

	fmt.Println("slice1 is ", slice1)
	fmt.Println("slice2 is ", slice2)
	fmt.Println("slice3 is ", slice3)
	fmt.Println("slice4 is ", slice4)
}
