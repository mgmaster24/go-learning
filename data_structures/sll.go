package gol_datastructures

import "fmt"

type SllNode[T comparable] struct {
	next *SllNode[T]
	val  T
}

type SinglyLinkedList[T comparable] struct {
	head *SllNode[T]
}

func (list *SinglyLinkedList[T]) PushBack(val T) {
	if list.head == nil {
		list.head = &SllNode[T]{
			next: nil,
			val:  val,
		}
	} else {
		currNode := list.head
		for currNode.next != nil {
			currNode = currNode.next
		}

		currNode.next = &SllNode[T]{
			next: nil,
			val:  val,
		}
	}
}

func (list *SinglyLinkedList[T]) PushFront(val T) {
	newHead := &SllNode[T]{
		next: list.head,
		val:  val,
	}

	list.head = newHead
}

func (list *SinglyLinkedList[T]) InsertBefore(beforeVal, val T) {
	found := list.Find(beforeVal)
	if found != nil {
		newNode := &SllNode[T]{
			val:  found.val,
			next: found.next,
		}

		found.val = val
		found.next = newNode
	}
}

func (list *SinglyLinkedList[T]) InsertAfter(afterVal, val T) {
	found := list.Find(afterVal)
	if found != nil {
		newNode := &SllNode[T]{
			val:  val,
			next: found.next,
		}

		found.next = newNode
	}
}

func (list *SinglyLinkedList[T]) Delete(val T) {
	currentNode := list.head
	var previousNode *SllNode[T] = nil
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

func (list *SinglyLinkedList[T]) Find(val T) *SllNode[T] {
	currentNode := list.head
	for currentNode != nil {
		if currentNode.val == val {
			return currentNode
		}

		currentNode = currentNode.next
	}

	return nil
}

func (list *SinglyLinkedList[T]) Print() {
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
