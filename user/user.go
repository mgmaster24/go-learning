package gol_user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("User expects valid input ")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u User) OutputUserData() {
	fmt.Printf("Hello %s %s!\n", u.firstName, u.lastName)
	fmt.Printf("Your birthday is %s.\n", u.birthDate)
	fmt.Printf("User created on: %v\n", u.createdAt.Local().String())
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
