package main

import (
	"fmt"
)

func main() {
	// arrExercise()
	mapsExercise()
	// id, title, price := getProductData()
}

func mapsExercise() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites["Amazon Web Services"])
	websites["LinkedIn"] = "https://linkedin.com"

	for key, value := range websites {
		fmt.Printf("Key: %v, Value: %v\n", key, value)
	}
}

func arrExercise() {
	arrayOfHobbies := [3]string{"Guitar", "Running", "Swimming"}
	fmt.Println(arrayOfHobbies)
	first := arrayOfHobbies[:1]
	fmt.Println(first)
	second := arrayOfHobbies[1:3]
	fmt.Println(second)

	slice1 := first[0:2]
	fmt.Println(slice1)

	slice2 := slice1[1:3]
	fmt.Println(slice2)

	dynGoals := []string{"Golang", "Gin"}
	fmt.Println("Before : ", dynGoals)
	dynGoals[1] = "GRpc"
	dynGoals = append(dynGoals, "Kubernetes")
	fmt.Println("After : ", dynGoals)
}

// func getProductInput(prompt string) string {
// 	fmt.Println(prompt)
// 	reader := bufio.NewReader(os.Stdin)
// 	text, err := reader.ReadString('\n')
// 	if err != nil {
// 		return ""
// 	}

// 	text = strings.TrimSpace(text)
// 	return text
// }

// func getProductData() (string, string, string) {
// 	id := getProductInput("Enter ID : ")
// 	title := getProductInput("Enter Title")
// 	price := getProductInput("Enter price")
// 	return id, title, price
// }
