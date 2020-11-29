package linkedlist

import (
	"fmt"
	"os"
	"testing"
)

var node *ListNode

func TestMain(m *testing.M) {
	node = &ListNode{Data: 1}
	node.Next = &ListNode{Data: 2}
	node.Next.Next = &ListNode{Data: 3}
	os.Exit(m.Run())
}

func TestListNode_Print(t *testing.T) {
	Print(node)
}

// AddByEmail a new node to end of the linked list
func TestListNode_Add(t *testing.T) {
	Add(&node, 1)
	Print(node)
}

func TestListNode_AddFront(t *testing.T) {
	addFront(&node, 1)
	Print(node)
}

// remove a list node of certain index in the linked list
func TestListNode_Remove(t *testing.T) {
	t.Run("remove the first element", func(t *testing.T) {
		remove(&node, 0)
		Print(node)
	})

	t.Run("remove other elements", func(t *testing.T) {
		remove(&node, 1)
		remove(&node, 1)
		Print(node)
	})
}

func TestListNode_Size(t *testing.T) {
	fmt.Print(Size(node))
}

// get the value of the linked list
func TestListNode_Get(t *testing.T) {
	fmt.Print(get(node, 1))
}
