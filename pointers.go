package main

import "fmt"

func main() {
	fmt.Println()
	age := 26
	agePointer := &age
	// & -> allows user to fetch the memory address of the declared age variable

	fmt.Println("Age : ", age, "Pointer : ", agePointer)
	fmt.Println()
	// Age :  26 Pointer :  0x14000112018 <- memory address

	// to derive the value of the variable from its memory address we can use
	// de-reference (*) symbol with the pointer variable

	fmt.Println("Deference Example")
	// for e.g
	fmt.Printf("De-reference %v\n\n", *agePointer)

	// adultYears := getAdultYears(agePointer)
	getAdultYears(agePointer)
	fmt.Println("Adult years after mutation", age)
}

func getAdultYears(age *int) {
	*age = *age - 18 //here we are overwriting the age's value at its memory address
	// return *age - 18
}

/* IMPORTANT NOTE:

fmt.Scan() uses a memory address (Pointer) as argument to take input from the user,
hence we use '&VARIABLE' as a parameter in scan to de-reference it

All values in Go have a so-called "Null Value" - i.e.,
the value that's set as a default if no value is assigned to a variable.

For example, the null value of an int variable is 0. Of a float64,
it would be 0.0. Of a string, it's "".

For a pointer, it's nil - a special value built-into Go.
nil represents the absence of an address value - i.e.,
a pointer pointing at no address / no value in memory.
*/
