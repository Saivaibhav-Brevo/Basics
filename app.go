package main

import "fmt"

func main() {

	var balance float64 = 0.0

	fmt.Println("Hello Welcome to Bank!")

	fmt.Println("Enter you choice:")

	fmt.Println(`1. Deposit 2. WWithdraw 3. Check Balance 4. Exit `)

	var choice int
	fmt.Scan(&choice)

	fmt.Println("Your selected choice is:", choice)

	if choice == 1 {
		balance = deposit(balance)
		fmt.Println("Your current balance is:", balance)
	} else if choice == 2 {
		balance = withdraw(balance)
		fmt.Println("Your current balance is:", balance)
	} else if choice == 3 {
		checkBalance(balance)
	} else {
		fmt.Println("Thank you for using our services!")
	}

}

func deposit(balance float64) float64 {
	fmt.Println("Deposit function called")
	var amount float64

	fmt.Println("Enter amount to be deposited")

	fmt.Scan(&amount)

	balance += amount

	return balance
}

func withdraw(balance float64) float64 {
	fmt.Println("Deposit function called")
	var amount float64

	fmt.Println("Enter amount to be withdrawn")

	fmt.Scan(&amount)

	balance -= amount

	return balance
}

func checkBalance(balance float64) {
	fmt.Println("Your current balance is:", balance)
}
