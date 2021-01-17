package sort

import (
	"algorithms-go/datastructures"
)

// BST performs BST-AVL sort
func BST(lista []int) []int {
	head := datastructures.NewNode(lista[0])
	tree := datastructures.Tree{Head: &head}
	for i, v := range lista {
		if i != 0 {
			presentNode := datastructures.NewNode(v)
			datastructures.InsertAVL(&presentNode, tree.Head, &tree)
		}
	}
	sortedList := make([]int, len(lista))
	sortedList, _ = datastructures.DepthFristTraverse(tree.Head, sortedList, 0, "in")
	return sortedList
}
