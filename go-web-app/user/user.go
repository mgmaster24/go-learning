package user

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Age         int       `json:"age"`
	BirthDate   time.Time `json:"birth_data"`
}

func NewUser() *User {
	return &User{}
}

func NewUserFromParts(firstName, lastName, phoneNumber, birthDate string, age int) (User, error) {
	date, err := time.Parse("2006-01-02", strings.TrimSpace(birthDate))
	if err != nil {
		return User{}, err
	}

	return User{
		FirstName:   strings.TrimSpace(firstName),
		LastName:    strings.TrimSpace(lastName),
		PhoneNumber: strings.TrimSpace(phoneNumber),
		Age:         age,
		BirthDate:   date,
	}, nil
}

func (u *User) PrintFormatted() {
	year, month, day := u.BirthDate.Date()
	fmt.Println("First Name:", u.FirstName)
	fmt.Println("Last Name:", u.LastName)
	fmt.Println("Phone:", u.PhoneNumber)
	fmt.Println("Age:", u.Age)
	fmt.Println("Birth Date:", fmt.Sprintf("%s %d, %d", month, day, year))
}

func (u *User) GetFromStdin() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("First Name: ")
	firstName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	u.FirstName = strings.TrimSpace(firstName)

	fmt.Print("Last Name: ")
	lastName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	u.LastName = strings.TrimSpace(lastName)

	fmt.Print("Phone: ")
	phoneNumber, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	u.PhoneNumber = strings.TrimSpace(phoneNumber)

	fmt.Print("Age: ")
	age, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	u.Age, err = strconv.Atoi(strings.TrimSpace(age))
	if err != nil {
		return err
	}

	fmt.Print("Birth Date: ")
	birthDate, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	date, err := time.Parse("2006-01-02", strings.TrimSpace(birthDate))
	if err != nil {
		return err
	}

	u.BirthDate = date
	return err
}
