package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mgmaster24/go-web-app/config"
)

func addValues(x, y int) int {
	return x + y
}

func divide(w http.ResponseWriter, r *http.Request) {
	var x, y float32 = 100.0, 10.0
	res, err := divVals(x, y)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", x, y, res)
}

func divVals(x, y float32) (float32, error) {
	if y <= 0 {
		return 0, errors.New("Division by zero!")
	}

	return x / y, nil
}

func main() {
	appConfig := config.NewAppConfig(8080)
	// color := "Green"
	// log.Println("Current color is:", color)
	// changeViaPointer(&color)
	// log.Println("Color is now:", color)

	// myMap2 := make(map[string]string)
	// myMap2["dog"] = "Bruno"
	// log.Println(myMap2["dog"])
	//
	// var users []user.BaseUser
	//
	// for {
	// 	fmt.Print("Add User? (q to quit): ")
	// 	var ch string
	// 	fmt.Scanln(&ch)
	// 	if ch == "q" {
	// 		break
	// 	}
	//
	// 	user := user.NewUser()
	// 	err := user.GetFromStdin()
	// 	if err != nil {
	// 		log.Panic("Get user failed", err)
	// 	}
	// 	users = append(users, user)
	// }
	//
	// fmt.Println("\nAdded Users:")
	// for _, user := range users {
	// 	user.PrintFormatted()
	// }

	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/divide", divide)
	fmt.Println("Starting web server...")
	http.ListenAndServe(appConfig.GetPortString(), nil)
}

//func changeViaPointer(s *string) {
//	*s = "Red"
//}
