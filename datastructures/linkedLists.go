package datastructures

import "fmt"

// NodeSLL is the basic building block for singly linked list.
type NodeSLL struct {
	value int
	Next  *NodeSLL
}

// SinglyLinkedList contains a pointer to the head of a SL-list.
type SinglyLinkedList struct {
	Head *NodeSLL
}

// CreateSLL is a utility function to create a DLL
func CreateSLL(value int) NodeSLL {
	node := NodeSLL{value: value, Next: nil}
	return node
}

// AppendSLL adds element to singly linked list
func AppendSLL(list *SinglyLinkedList, AddedNode *NodeSLL) {
	head := list.Head
	appendSLL(head, AddedNode)
}

func appendSLL(PresentNode *NodeSLL, AddedNode *NodeSLL) {
	if PresentNode.Next != nil {
		appendSLL(PresentNode.Next, AddedNode)
		return
	}
	PresentNode.Next = AddedNode
}

// TraverseSLL traverses a single linked list in a forward looking fashion.
func TraverseSLL(List *SinglyLinkedList) {
	traverseSLL(List.Head)
}

func traverseSLL(PresentNode *NodeSLL) {
	fmt.Println(PresentNode.value)

	// If node has child recurse into it.
	if PresentNode.Next != nil {
		traverseSLL(PresentNode.Next)
	}
}

// FindNodeSLL finds and returns, if possible, a node of a given value.
func FindNodeSLL(value int, List *SinglyLinkedList) NodeSLL {
	node, _ := findNodeSLL(value, List.Head)
	return node
}

func findNodeSLL(value int, PresentNode *NodeSLL) (NodeSLL, bool) {
	if value == PresentNode.value {
		return *PresentNode, true
	}
	if PresentNode.Next != nil {
		return findNodeSLL(value, PresentNode.Next)
	}
	return *PresentNode, false
}

// RemoveFromSLL removes a node with a given value from a singly linked list.
func RemoveFromSLL(value int, list *SinglyLinkedList) {
	removedHead := removeFromSLL(value, list.Head)
	if removedHead {
		list.Head = list.Head.Next
	}
}

func removeFromSLL(value int, PresentNode *NodeSLL) bool {
	if PresentNode.value == value {
		return true
	}
	if PresentNode.Next != nil {
		if PresentNode.Next.value == value {
			PresentNode.Next = PresentNode.Next.Next
			return false
		}
		return removeFromSLL(value, PresentNode.Next)
	}
	return false
}

// NodeDLL is the basic building block for double linked list
type NodeDLL struct {
	Value int
	Next  *NodeDLL
	Past  *NodeDLL
}

// DoubleLinkedList contains a pointer to the head and tail of a DL-list
type DoubleLinkedList struct {
	Head *NodeDLL
	Tail *NodeDLL
}

// CreateDLL is a utility function to create a DLL
func CreateDLL(node *NodeDLL) DoubleLinkedList {
	list := DoubleLinkedList{Head: node, Tail: node}
	return list
}

// TraverseDLL traverses linked list going forwards.
func TraverseDLL(list DoubleLinkedList, direction string) {
	directionToFunc := map[string]func(*NodeDLL){
		"forwards":  traverseForwardsDLL,
		"backwards": traverseBackwardsDLL,
	}
	directionToEnd := map[string]*NodeDLL{
		"forwards":  list.Head,
		"backwards": list.Tail,
	}
	directionToFunc[direction](directionToEnd[direction])
}

func traverseForwardsDLL(presentNode *NodeDLL) {
	fmt.Println(presentNode.Value)
	if presentNode.Next != nil {
		traverseForwardsDLL(presentNode.Next)
	}
}

func traverseBackwardsDLL(presentNode *NodeDLL) {
	fmt.Println(presentNode.Value)
	if presentNode.Past != nil {
		traverseBackwardsDLL(presentNode.Past)
	}
}

// AppendDLL adds element to the end of a double linked list
func AppendDLL(list *DoubleLinkedList, insertedNode *NodeDLL) {
	presentNode := list.Head
	appendDLL(presentNode, insertedNode)
	list.Tail = insertedNode
}

func appendDLL(presentNode *NodeDLL, insertedNode *NodeDLL) {
	if presentNode.Next != nil {
		appendDLL(presentNode.Next, insertedNode)
	} else {
		presentNode.Next = insertedNode
		insertedNode.Past = presentNode
	}
}

// InsertDLL adds element to the begining of a double linked list
func InsertDLL(list *DoubleLinkedList, insertedNode *NodeDLL) {
	presentNode := list.Head
	insertedNode.Next = presentNode
	presentNode.Past = insertedNode
	list.Head = insertedNode

}

// RemoveFromDLL removes node of given value if found in linked list
func RemoveFromDLL(Value int, list *DoubleLinkedList) {
	presentNode := list.Head
	foundNode, validation := removeDLL(Value, presentNode)
	if validation[2] == false {
		list.Head = foundNode.Next
	}
	if validation[1] == false {
		list.Tail = foundNode.Past
	}

}
func removeDLL(Value int, PresentNode *NodeDLL) (NodeDLL, [3]bool) {
	foundNode, err := findNodeDLL(Value, PresentNode)
	validation := [3]bool{false, false, false}
	if err {
		return foundNode, validation
	}
	validation[0] = true
	if foundNode.Next != nil {
		validation[1] = true
		foundNode.Next.Past = foundNode.Past
	}
	if foundNode.Past != nil {
		validation[2] = true
		foundNode.Past.Next = foundNode.Next
	}
	return foundNode, validation

}

// findNodeDLL is a utility function used to check if a node existis
// and return it if possible
func findNodeDLL(Value int, presentNode *NodeDLL) (NodeDLL, bool) {
	if presentNode.Value == Value {
		return *presentNode, false
	}
	if presentNode.Next == nil {
		return *presentNode, true
	}
	return findNodeDLL(Value, presentNode.Next)
}
