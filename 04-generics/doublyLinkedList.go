package main

import "fmt"

type Order interface {
	int | float64 | string
}

type item[T any] struct {
	Data T
	prev *item[T]
	next *item[T]
}

type doublyLinkedList[T Order] struct {
	head *item[T]
	tail *item[T]
}

func (dll *doublyLinkedList[T]) Add(data T) {
	newItem := &item[T]{
		Data: data,
		prev: nil,
		next: nil,
	}

	if dll.head == nil {
		dll.head = newItem
		dll.tail = newItem
		return
	}

	newItem.prev = dll.tail
	dll.tail.next = newItem
	dll.tail = newItem
}

func (dll *doublyLinkedList[T]) Delete(data T) bool {
	// 链表为空
	if dll.head == nil {
		return false
	}

	current := dll.head
	for current != nil {
		if current.Data == data {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				dll.head = current.next
			}

			if current.next != nil {
				current.next.prev = current.prev
			} else {
				dll.tail = current.prev
			}
			return true
		}
		current = current.next
	}
	return false
}

func (dll *doublyLinkedList[T]) PrintMe() {
	current := dll.head
	for current != nil {
		fmt.Printf("%v -> ", current.Data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var dll doublyLinkedList[int]

	// Add elements
	dll.Add(10)
	dll.Add(20)
	dll.Add(30)
	dll.Add(40)
	dll.PrintMe() // 10 -> 20 -> 30 -> 40 -> nil

	dll.Delete(10)
	dll.Delete(30)
	dll.PrintMe() // 20 -> 40 -> nil
}
