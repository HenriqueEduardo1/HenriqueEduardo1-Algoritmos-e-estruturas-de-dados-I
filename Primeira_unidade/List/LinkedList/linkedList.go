package main

import (
	"fmt"
	"strings"
)

type List[T any] interface {
	Show()
	Append(value T)
	Size() int
	Get(index int) T
	Pop() T
	Update(value T, index int)
	Insert(value T, index int)
	Remove(index int)
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head     *Node[T]
	inserted int
}

func (a *LinkedList[T]) Size() int {
	return a.inserted
}

func (a *LinkedList[T]) Show() {
	var sb strings.Builder

	sb.WriteString("(")
	current := a.head

	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.value))

		if current.next != nil {
			sb.WriteString(", ")
		}

		current = current.next
	}

	sb.WriteString(")")

	fmt.Println(sb.String())
}

func (a *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{value: value}

	if a.head == nil {
		a.head = newNode
	} else {
		current := a.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
		a.inserted++
	}
}

func (a *LinkedList[T]) Get(index int) T {

	if index < 0 || index >= a.inserted {
		fmt.Println("Index out of bounds")
	} else if a.head == nil {
		fmt.Println("Empty list")
	} else {
		current := a.head

		for i := 0; i < index; i++ {
			current = current.next
		}

		return current.value
	}

	var defaultValue T
	return defaultValue
}

func main() {
	list := LinkedList[int]{head: nil, inserted: 0}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	fmt.Println(list.Get(-1))
	fmt.Println(list.Get(5))
}
