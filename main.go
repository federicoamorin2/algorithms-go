package main

import (
	"algorithms-go/sort"
	"fmt"
	"math/rand"
	"time"
)

func testAlgoSpeed() {
	for i := 2; i < 2000; i += 50 {
		a := make([]int, i)
		for j := 0; j < i; j++ {
			a[j] = rand.Int()
		}
		start := time.Now()
		a = sort.MaxHeapSort(a)
		fmt.Println(a)
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(len(a))
		fmt.Println(elapsed)
	}
}
func main() {

	a := []int{8, 23, 5, 12, 2, 1, 22}

	b := sort.MergeSort(a, false)
	c := sort.MaxHeapSort(a)
	fmt.Println(b)
	fmt.Println(c)
	// testAlgoSpeed()
}
