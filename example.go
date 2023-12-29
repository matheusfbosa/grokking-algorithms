package main

import (
	"fmt"

	"github.com/matheusfbosa/grokking-algorithms/chapter1"
	"github.com/matheusfbosa/grokking-algorithms/chapter2"
	"github.com/matheusfbosa/grokking-algorithms/chapter3"
)

func main() {
	//examplesChapter1()
	//examplesChapter2()
	examplesChapter3()
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

func examplesChapter2() {
	// Selection Sort
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("[Selection Sort] Unsorted array:", arr)
	sorted := chapter2.SelectionSort(arr)
	fmt.Println("[Selection Sort] Sorted array:", sorted)

	// Linked List
	linkedList := chapter2.NewLinkedList()
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	fmt.Print("[Linked List] ")
	linkedList.Display()
}

func examplesChapter3() {
	// Recursion
	fmt.Printf("[Recursion] %d! = %d\n", 5, chapter3.Factorial(5))
}
