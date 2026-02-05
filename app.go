package main

import "fmt"

func main() {
	fmt.Println("Hello Welcome to Bank!")

	fmt.Println("Enter you choice:")

	fmt.Println(`1. Deposit 2. WWithdraw 3. Check Balance 4. Exit `)

	var choice int
	fmt.Scan(&choice)

	fmt.Println("Your selected choice is:", choice)

	if choice == 1 {
		deposit()
	}

}

func deposit() {
	fmt.Println("Deposit function called")
}
