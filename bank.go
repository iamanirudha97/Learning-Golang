package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}

	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse/write to balance.txt file")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceStr := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceStr), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("----------------------------")
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("----------------------------")
	}

	fmt.Println("Welcome to the simple Go Banking Application")

	for {
		fmt.Println("1. Check Balance: ")
		fmt.Println("2. Deposit Money: ")
		fmt.Println("3. Withdraw Money: ")
		fmt.Println("4. Exit ")

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
				writeBalanceToFile(accountBalance)
			}

		case 3:
			fmt.Println("Your current balance is : ", accountBalance)
			fmt.Print("Enter the amount to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid Amount, deposit must be greater than 0. ")
				continue
				// return
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds to withdraw, withdrawal must be withing your available balance. ")
			} else {
				accountBalance -= withdrawalAmount
				fmt.Println("New balance after Withdrawal is: ", accountBalance)
				writeBalanceToFile(accountBalance)
			}

		default:
			fmt.Println("Logging out")
			fmt.Println("Thanks for Banking with us")
			return
		}
	}
}
