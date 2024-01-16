package gol_exercises

import (
	"fmt"

	gol_io "go-learning.com/learning/io"
	gol_user "go-learning.com/learning/user"
)

func RunStructsExercise() {
	firstName := getUserData("Please enter your firstName: ")
	lastName := getUserData("Please enter your lastName: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user, err := gol_user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := gol_user.NewAdmin("admin@golearning.com", "ASDSAFASDDFASD")

	user.OutputUserData()
	user.ClearUserName()
	user.OutputUserData()
	admin.OutputUserData()
}

func getUserData(context string) (val string) {
	gol_io.ScanForInput(context, &val)
	return
}
