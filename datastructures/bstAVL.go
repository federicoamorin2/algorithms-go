package datastructures

// Node defines basic properties for nodes of the BTS structure.
type Node struct {
	parent     *Node
	leftChild  *Node
	rightChild *Node
	value      int
	height     int
	balance    int
}

// Tree keeps track of Tree head.
type Tree struct {
	Head *Node
}

// NewNode is a intialization function for nodes
func NewNode(value int) Node {
	var newNode Node
	newNode.value = value
	newNode.height = 0
	newNode.balance = 0
	return newNode
}

// InsertAVL performs BST insertion using AVL balancing.
func InsertAVL(nodeInserted *Node, startNode *Node, tree *Tree) {
	// Check if node is smaller than or equal initial node.
	if startNode.value >= nodeInserted.value {
		// If node has left child recurse to insert in child node
		if startNode.leftChild != nil {
			InsertAVL(nodeInserted, startNode.leftChild, tree)
		} else {
			// If not insert in present node
			startNode.leftChild = nodeInserted
			nodeInserted.parent = startNode
		}
	} else {
		// If node has right child recurse to insert in child node
		if startNode.rightChild != nil {
			InsertAVL(nodeInserted, startNode.rightChild, tree)
		} else {
			// If not insert in present node
			startNode.rightChild = nodeInserted
			nodeInserted.parent = startNode
		}
	}
	computeBalanceFactor(startNode)
	// Balance the tree
	balanceTree(startNode, tree)
	return
}

func balanceTree(startNode *Node, tree *Tree) {

	if startNode.balance > 1 {

		if startNode.leftChild.balance > 0 {
			if startNode.parent == nil {
				leftRotate(startNode)
				tree.Head = startNode.parent
			} else {
				leftRotate(startNode)
			}
		} else if startNode.leftChild.balance < 0 {
			rightRotate(startNode.leftChild)
			if startNode.parent == nil {
				leftRotate(startNode)
				tree.Head = startNode.parent
			} else {
				leftRotate(startNode)
			}

		}
	} else if startNode.balance < -1 {
		if startNode.rightChild.balance > 0 {
			leftRotate(startNode.rightChild)
			if startNode.parent == nil {
				rightRotate(startNode)
				tree.Head = startNode.parent
			} else {
				rightRotate(startNode)
			}
		} else if startNode.rightChild.balance < 0 {
			if startNode.parent == nil {
				rightRotate(startNode)
				tree.Head = startNode.parent
			} else {
				rightRotate(startNode)
			}
		}
	}
}

// leftRotate does a right rotation, as seen in drawing below
func leftRotate(presentNode *Node) {
	//        --a
	//       |
	//   -- b     ==>    -- b --
	//  |               |       |
	//  c               c       a
	A := presentNode
	B := presentNode.leftChild
	if A.parent != nil {
		if A == A.parent.leftChild {
			A.parent.leftChild = B
		} else {
			A.parent.rightChild = B
		}
	}
	B.parent = A.parent
	A.parent = B
	A.leftChild = B.rightChild
	if B.rightChild != nil {
		B.rightChild.parent = A
	}
	B.rightChild = A
	computeBalanceFactor(A)
	computeBalanceFactor(B)
}

// rightRotate does a right rotation, as seen in drawing below
func rightRotate(presentNode *Node) {
	//  a--
	//     |
	//     b --    ==>     -- b --
	//         |          |       |
	//         c          a       c
	A := presentNode
	B := presentNode.rightChild
	if A.parent != nil {
		if A == A.parent.leftChild {
			A.parent.leftChild = B
		} else {
			A.parent.rightChild = B
		}
	}
	B.parent = A.parent
	A.parent = B
	A.rightChild = B.leftChild
	if B.leftChild != nil {
		B.leftChild.parent = A
	}
	B.leftChild = A
	computeBalanceFactor(A)
	computeBalanceFactor(B)
}

// Adjust the balance factor with new insertion
func computeBalanceFactor(presentNode *Node) {
	if presentNode.rightChild != nil && presentNode.leftChild != nil {
		presentNode.balance = presentNode.leftChild.height - presentNode.rightChild.height
		presentNode.height = Max(presentNode.leftChild.height, presentNode.rightChild.height) + 1
	} else if presentNode.rightChild == nil && presentNode.leftChild != nil {
		presentNode.height = presentNode.leftChild.height + 1
		presentNode.balance = presentNode.height
	} else if presentNode.rightChild != nil && presentNode.leftChild == nil {
		presentNode.height = presentNode.rightChild.height + 1
		presentNode.balance = -presentNode.height
	} else {
		presentNode.height = 0
		presentNode.balance = 0
	}
	return
}

// Abs returns the biggest of two integer numbers
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Max returns the biggest of two integer numbers
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func findNode(value int, presentNode *Node) *Node {
	// If value is equal to present node value return present node.
	if presentNode.value == value {
		return presentNode
		// If value is greater than present node value and present node has
		// a right child recurse into right child.
	} else if presentNode.value >= value && presentNode.rightChild != nil {
		return findNode(value, presentNode.rightChild)
		// Same for right side.
	} else if presentNode.leftChild != nil {
		return findNode(value, presentNode.leftChild)
	}
	// If all cases above fail node was not found and nil is returned
	return nil

}

func deleteNode(value int, startNode *Node, tree *Tree) string {
	// Find node in tree if it doesn't exist return not found.
	foundNode := findNode(value, startNode)
	if foundNode == nil {
		return "Not found"
	}

	// Get parent node and see wether found node is left or right child.
	parentNode := foundNode.parent
	if parentNode.leftChild == foundNode {
		parentNode.leftChild = nil
	} else {
		parentNode.rightChild = nil
	}

	// If found node has child nodes re-insert them in tree
	if foundNode.leftChild != nil {
		InsertAVL(foundNode.leftChild, startNode, tree)
	}
	if foundNode.rightChild != nil {
		InsertAVL(foundNode.rightChild, startNode, tree)
	}
	return "Removed"
}

func findMax(presentNode *Node) *Node {
	// If present node doesn't have right child return node else recurse
	// into right child.
	if presentNode.rightChild == nil {
		return presentNode
	}
	return findMax(presentNode.rightChild)

}

func findMin(presentNode *Node) *Node {
	// If present node doesn't have left child return node else recurse
	// into left child.
	if presentNode.leftChild == nil {
		return presentNode
	}
	return findMin(presentNode.leftChild)

}

// DepthFristTraverse traverses tree, order depends on order argument.
func DepthFristTraverse(presentNode *Node, lista []int, listIndex int, order string) ([]int, int) {
	if order == "pre" {
		lista[listIndex] = presentNode.value
		listIndex++
	}
	if presentNode.leftChild != nil {
		lista, listIndex = DepthFristTraverse(presentNode.leftChild, lista, listIndex, order)
	}
	if order == "in" {
		lista[listIndex] = presentNode.value
		listIndex++
	}
	if presentNode.rightChild != nil {
		lista, listIndex = DepthFristTraverse(presentNode.rightChild, lista, listIndex, order)
	}
	if order == "post" {
		lista[listIndex] = presentNode.value
		listIndex++
	}
	return lista, listIndex
}
