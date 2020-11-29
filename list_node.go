package linkedlist

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

func New(data ...int) *ListNode {
	front := &ListNode{Data: data[0]}
	for _, d := range data[1:] {
		Add(&front, d)
	}

	return front
}

func AddNode(front **ListNode, node *ListNode) {
	if front == nil {
		front = &node
		return
	}
	current := *front
	for current.Next != nil {
		current = current.Next
	}
	current.Next = node
}

func Add(front **ListNode, data int) {
	newNode := &ListNode{Data: data}
	AddNode(newNode)
}

func addFront(front **ListNode, data int) {
	*front = &ListNode{Data: data, Next: *front}
}

func remove(front **ListNode, index int) {
	count := Size(*front)
	if index > count-1 || index < 0 {
		panic("index out of range")
	}

	if front == nil {
		return
	}

	if index == 0 {
		*front = (*front).Next
		return
	}

	current := *front

	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	current.Next = current.Next.Next
}

func Size(front *ListNode) int {
	current := front
	count := 0
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func get(front *ListNode, index int) int {
	if index > Size(front)-1 || index < 0 {
		panic("index out of range")
	}
	current := front
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Data
}

func Print(front *ListNode) {
	current := front
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
}
