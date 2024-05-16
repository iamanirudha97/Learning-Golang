package main

import "fmt"

func main() {
	var listOfStrings [4]string = [4]string{"Anirudha", "Sharvil", "Daryl", "Mohit"}
	pieces := listOfStrings[:3]
	fmt.Println(pieces)
}
