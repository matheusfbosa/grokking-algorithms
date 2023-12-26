package main

import (
	"fmt"

	"github.com/matheusfbosa/grokking-algorithms/chapter1"
)

func main() {
	examplesChapter1()
}

func examplesChapter1() {
	// Binary Search
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bsTarget := 5
	bsResult := chapter1.BinarySearch(arr, bsTarget)
	if bsResult != -1 {
		fmt.Printf("[Binary Search] Element %d found at index %d in the list %v\n", bsTarget, bsResult, arr)
	}

	// Binary Search Tree (BST)
	var root *chapter1.Node
	keys := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, key := range keys {
		root = root.Insert(key)
	}
	fmt.Printf("[BST] Tree in-order traversal: ")
	root.InOrderTraversal()
	targetVal := 6
	bstResult := root.Search(targetVal)
	if bstResult != nil {
		fmt.Printf("\n[BST] Element %d found in the tree\n", targetVal)
	}
}
