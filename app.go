package main

import (
	"fmt"
	"os"
)

func main() {

	var balance float64 = 0.0

	fmt.Println("Hello Welcome to Bank!")

	var choice int

	for choice != 4 {

		fmt.Println("--------------------------------")
		fmt.Println("Enter you choice:")

		fmt.Println(`1. Deposit 2. WWithdraw 3. Check Balance 4. Exit `)

		fmt.Scan(&choice)

		fmt.Println("Your selected choice is:", choice)

		fmt.Println("--------------------------------")

		switch choice {
		case 1:
			balance = deposit(balance)
		case 2:
			balance = withdraw(balance)
		case 3:
			checkBalance(balance)
		case 4:
			return
		default:
			return
		}
	}

}

func deposit(balance float64) float64 {

	var amount float64

	fmt.Println("Enter amount to be deposited")

	fmt.Scan(&amount)

	balance += amount

	return balance
}

func withdraw(balance float64) float64 {

	var amount float64

	fmt.Println("Enter amount to be withdrawn")

	fmt.Scan(&amount)

	if amount > balance {
		fmt.Println("Insufficient balance!")
		return balance
	}

	balance -= amount

	return balance
}

func checkBalance(balance float64) {
	fmt.Println("Your current balance is:", balance)
	writeBalanceToFile(balance)
}

func writeBalanceToFile(balance float64) {

	balanceStr := fmt.Sprintf("Current Balance: %.2f\n", balance)

	os.WriteFile("hello.txt", []byte(balanceStr), 0644)

}
