package main

import "fmt"

func main() {

	var arr [5]int

	fmt.Println("Enter array elements:")

	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("arr is ", arr)

	temp := arr[0]
	for i := 1; i < 5; i++ {
		if temp < arr[i] {
			temp = arr[i]
		}
	}
	fmt.Println("Largest element is ", temp)
}
