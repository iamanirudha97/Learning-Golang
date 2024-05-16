package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(fileName)
	if err != nil {
		fmt.Println("----------------------------")
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("----------------------------")
	}

	fmt.Println("Welcome to the simple Go Banking Application")
	fmt.Println("Contact Us : ", randomdata.PhoneNumber())
	for {
		userOptions()
		var choice int
		fmt.Print(" Enter your action: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accountBalance)

		case 2:
			fmt.Println("Your current balance is : ", accountBalance)
			fmt.Print("Enter the amount to deposit: ")
			var userDeposit float64
			fmt.Scan(&userDeposit)
			if userDeposit <= 0 {
				fmt.Println("Invalid Amount, deposit must be greater than 0. ")
				continue
				// return
			} else {
				accountBalance += userDeposit
				fmt.Println("Your Balance after deposit is: ", accountBalance)
				fileops.WriteFloatToFile(accountBalance, fileName)
			}

		case 3:
			fmt.Println("Your current balance is : ", accountBalance)
			fmt.Print("Enter the amount to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount, deposit must be greater than 0. ")
				continue
				// return will result in exiting to the terminal
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds to withdraw, withdrawal must be withing your available balance. ")
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("New balance after Withdrawal is: ", accountBalance)
				fileops.WriteFloatToFile(accountBalance, fileName)
			}

		default:
			fmt.Println("Logging out")
			fmt.Println("Thanks for Banking with us")
			return
		}
	}
}
