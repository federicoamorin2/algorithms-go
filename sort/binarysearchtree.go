package sort

import (
	"algorithms-go/datastructures"
)

// BST performs BST-AVL sort
func BST(lista []float32) []float32 {
	head := datastructures.NewNode(lista[0])
	tree := datastructures.Tree{Head: &head}
	for i, v := range lista {
		if i != 0 {
			presentNode := datastructures.NewNode(v)
			datastructures.InsertAVL(&presentNode, tree.Head, &tree)
		}
	}
	var lista2 []float32
	return datastructures.DepthFristTraverse(tree.Head, lista2, "in")
}
