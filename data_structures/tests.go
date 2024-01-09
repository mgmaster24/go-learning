package gol_datastructures

import "fmt"

func RunSLLTests() {
	var sll SinglyLinkedList[string]
	sll.PushBack("This")
	sll.PushBack("is")
	sll.PushBack("a test")
	sll.Print()
	fmt.Println()

	sll.PushFront("front")
	sll.PushFront("to")
	sll.PushFront("Nodes")
	sll.PushFront("Adding")
	sll.Print()
	fmt.Println()

	sll.InsertBefore("is", "has/")
	sll.Print()
	fmt.Println()

	sll.InsertAfter("a test", "OMG")
	sll.Print()
	fmt.Println()

	sll.Delete("Adding")
	sll.Delete("This")
	sll.Print()
	fmt.Println()
}
