package main

func main() {

	arr := []int{1, 2, 34, 56}

	slice1 := arr[1:3]

	slice1[0] = 45

	println("Array:", arr[0], arr[1], arr[2], arr[3])

}
