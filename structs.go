package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//Struct Literal OR Composite Literal
	// var person1 *person -> gives should merge decalration and assignment
	// var person1 *person = newPerson(userfirstName, userlastName, userbirthdate)

	// var person1 *user.Person
	// person1, err := user.New(userfirstName, userlastName, userbirthdate)
	// if err != nil {
	// 	fmt.Println()
	// 	panic(err)
	// }
	// person1.ShowOutput()
	// person1.ClearUserDetails()
	// person1.ShowOutput()

	fmt.Println("----------------------------------")

	admin := user.NewAdmin(userfirstName, userlastName, userbirthdate, "ok@ok.com", "password")
	admin.ShowOutput()
	admin.ClearUserDetails()
	admin.ShowOutput()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
