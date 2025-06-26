package main

import (
	"fmt"
)

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {

		fmt.Println("What'd you like to do today?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit app")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balanca updated! New balance:", accountBalance)
		case 3:
			fmt.Print("Withdrawal amount:")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("You have withdrawed:", withdrawalAmount)
			fmt.Println("Your new balance is:", accountBalance)
		default:
			fmt.Println("Have a good day!")
			fmt.Println("Thanks for chosing Go Bank!")
			return
		}

	}

}
