package datastructures

import "fmt"

// NodeSLL is the basic building block for singly linked list.
type NodeSLL struct {
	value int
	Next  *NodeSLL
}

// NodeDLL is the basic building block for double linked list
type NodeDLL struct {
	Value int
	Next  *NodeDLL
	Past  *NodeDLL
}

// DoubleLinkedList contains a pointer to the head of a DL-list
type DoubleLinkedList struct {
	head *NodeDLL
}

// TraverseDLL traverses linked list going forwards.
func TraverseDLL(list DoubleLinkedList) {
	presentNode := list.head
	traverseForwardsDLL(presentNode)
}

func traverseForwardsDLL(presentNode *NodeDLL) {
	fmt.Println(presentNode.Value)
	if presentNode.Next != nil {
		traverseForwardsDLL(presentNode.Next)
	}
}

func traverseBackwordsDLL(presentNode *NodeDLL) {
	fmt.Println(presentNode.Value)
	if presentNode.Past != nil {
		traverseBackwordsDLL(presentNode.Past)
	}
}

// AppendDLL adds element to the end of a double linked list
func AppendDLL(presentNode *NodeDLL, insertedNode *NodeDLL) {
	if presentNode.Next != nil {
		AppendDLL(presentNode.Next, insertedNode)
	} else {
		presentNode.Next = insertedNode
		insertedNode.Past = presentNode
	}
}

func insertDLL(presentNode *NodeDLL, insertedNode *NodeDLL) {
	if presentNode.Next != nil {
		insertDLL(presentNode.Next, insertedNode)
	} else {
		presentNode.Next = insertedNode
		insertedNode.Past = presentNode
	}
}

// RemoveDLL removes node of given value if found in linked list
func RemoveDLL(Value int, headNode *NodeDLL) string {
	foundNode, err := findNodeDLL(Value, headNode)
	if err {
		return "Node not found"
	}
	if foundNode.Past != nil {
		foundNode.Past.Next = foundNode.Next
	} else {
		// found node is head of list
		headNode.Next = nil
		headNode.Past = nil
		fmt.Println("cheguei aqui")
	}
	if foundNode.Next != nil {
		foundNode.Next.Past = foundNode.Past
	}
	foundNode.Next = nil
	foundNode.Past = nil
	return "Removed value"
}

func findNodeDLL(Value int, presentNode *NodeDLL) (NodeDLL, bool) {
	if presentNode.Value == Value {
		return *presentNode, false
	}
	if presentNode.Next == nil {
		return *presentNode, true
	}
	return findNodeDLL(Value, presentNode.Next)
}
