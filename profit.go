package main

import (
	"errors"
	"fmt"
	"os"
)

const file = "taxReturns.txt"

func writeToFile(result string) {
	os.WriteFile(file, []byte(result), 0644)
}

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("----------------------------")
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("----------------------------")
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println("----------------------------")
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("----------------------------")
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("----------------------------")
		fmt.Println("ERROR!!!")
		fmt.Println(err)
		fmt.Println("----------------------------")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	data := fmt.Sprint(ebt, profit, ratio)
	writeToFile(data)
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Invalid Input, input cant be zero")
	}

	return userInput, nil
}
