package main

import "fmt"

func main() {
	printData()
}

//recursion

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNUmbers := []int{}
	for _, value := range *numbers {
		dNUmbers = append(dNUmbers, transform(value))
	}
	return dNUmbers
}

func double(val int) int {
	return val * 2
}

func triple(val int) int {
	return val * 3
}

func printData() {
	fmt.Println()
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Before: ", numbers)
	doubleNums := transformNumbers(&numbers, double)
	tripleNums := transformNumbers(&numbers, triple)
	fmt.Println("Double: ", doubleNums)
	fmt.Println("Triple: ", tripleNums)
	fmt.Println("After: ", numbers)
	fmt.Println()
}
