package sort

import (
	"algorithms-go/datastructures"
)

// InsertArrayBST puts an array in a BST format.
func InsertArrayBST(lista []int) datastructures.Tree {
	head := datastructures.NewNode(lista[0])
	tree := datastructures.Tree{Head: &head}
	for i, v := range lista {
		if i != 0 {
			presentNode := datastructures.NewNode(v)
			datastructures.InsertAVL(&presentNode, tree.Head, &tree)
		}
	}
	return tree
}

// BstSort performs BST-AVL sort
func BstSort(lista []int) []int {
	tree := InsertArrayBST(lista)
	sortedList := make([]int, len(lista))
	sortedList, _ = datastructures.DepthFristTraverse(tree.Head, sortedList, 0, "in")
	return sortedList
}
