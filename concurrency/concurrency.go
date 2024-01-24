package gol_concurrency

import (
	"fmt"
	"time"
)

func Run() {
	//finished := make([]chan bool, 3)
	done := make(chan bool)

	// finished[0] = make(chan bool)
	// go saySomething("Hello", finished[0])
	// finished[1] = make(chan bool)
	// go respond("How are you", finished[1])
	// finished[2] = make(chan bool)
	// go saySomething("Nice to meet you", finished[2])

	go saySomething("Hello", done)
	go respond("How are you", done)
	go saySomething("Nice to meet you", done)

	// for _, finish := range finished {
	// 	<-finish
	// }

	for range done {
	}
}

func saySomething(someting string, done chan bool) {
	fmt.Println(someting)
	done <- true
}

func respond(resp string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(resp)
	done <- true
	close(done)
}
