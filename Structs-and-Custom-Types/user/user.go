package user

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	Person
}

func NewAdmin(firstName, lastName, birthdate, email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		Person: Person{
			firstName: firstName,
			lastName:  lastName,
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*Person, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name last name and birth date cant be empty")
	}
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

// Method Overiding
func (user *Person) ShowOutput() string {
	return fmt.Sprintf("%v %v %v", user.firstName, user.lastName, user.birthdate)
}

func (user *Admin) ShowOutput() {
	fmt.Printf("%v %v %v\n", user.Person.ShowOutput(), user.email, user.password)
}

func (user *Person) ClearUserDetails() {
	user.firstName = ""
	user.lastName = ""
}
