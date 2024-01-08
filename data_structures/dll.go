package gol_datastructures

import "fmt"

type DoublyLinkedList[T comparable] struct {
	head *DllNode[T]
}

func (list *DoublyLinkedList[T]) PushBack(val T) {
	if list.head == nil {
		list.head = &DllNode[T]{
			prev: nil,
			next: nil,
			val:  val,
		}
	} else {
		currNode := list.head
		for currNode.next != nil {
			currNode = currNode.next
		}

		currNode.next = &DllNode[T]{
			prev: currNode,
			next: nil,
			val:  val,
		}
	}
}

func (list *DoublyLinkedList[T]) PushFront(val T) {
	newHead := &DllNode[T]{
		prev: nil,
		next: list.head,
		val:  val,
	}

	list.head = newHead
}

func (list *DoublyLinkedList[T]) InsertBefore(beforeVal, val T) {
	found := list.Find(beforeVal)
	if found != nil {
		newNode := &DllNode[T]{
			val:  found.val,
			next: found.next,
		}

		found.val = val
		found.next = newNode
	}
}

func (list *DoublyLinkedList[T]) InsertAfter(afterVal, val T) {
	found := list.Find(afterVal)
	if found != nil {
		newNode := &DllNode[T]{
			val:  val,
			next: found.next,
		}

		found.next = newNode
	}
}

func (list *DoublyLinkedList[T]) Delete(val T) {
	currentNode := list.head
	var previousNode *DllNode[T] = nil
	for currentNode != nil {
		if currentNode.val == val {
			if previousNode == nil {
				list.head = list.head.next
			} else {
				previousNode.next = currentNode.next
			}
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}
}

func (list *DoublyLinkedList[T]) Find(val T) *DllNode[T] {
	currentNode := list.head
	for currentNode != nil {
		if currentNode.val == val {
			return currentNode
		}

		currentNode = currentNode.next
	}

	return nil
}

func (list *DoublyLinkedList[T]) Print() {
	if list == nil || list.head == nil {
		fmt.Println("Nothing to print")
	}

	curreNode := list.head
	for curreNode != nil {
		fmt.Print(curreNode.val)
		fmt.Print(" ")
		curreNode = curreNode.next
	}
}