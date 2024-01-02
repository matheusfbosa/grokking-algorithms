package main

import (
	"fmt"

	"github.com/matheusfbosa/grokking-algorithms/chapter1"
	"github.com/matheusfbosa/grokking-algorithms/chapter2"
	"github.com/matheusfbosa/grokking-algorithms/chapter3"
	"github.com/matheusfbosa/grokking-algorithms/chapter4"
	"github.com/matheusfbosa/grokking-algorithms/chapter5"
	"github.com/matheusfbosa/grokking-algorithms/helper"
)

func main() {
	//examplesChapter1()
	//examplesChapter2()
	//examplesChapter3()
	//examplesChapter4()
	examplesChapter5()
}

func examplesChapter1() {
	// Binary Search
	arr := helper.MakeRandomArray(5, 0, 10)
	bsTarget := 5
	bsResult := chapter1.BinarySearch(arr, bsTarget)
	if bsResult != -1 {
		fmt.Printf("[Binary Search] Element %d found at index %d in the list %v\n", bsTarget, bsResult, arr)
	}

	// Binary Search Tree (BST)
	var root *chapter1.Node
	keys := helper.MakeRandomArray(5, 0, 10)
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
	arr := helper.MakeRandomArray(10, -100, 100)
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

func examplesChapter4() {
	// Divide and Conquer
	arr := helper.MakeRandomArray(10, 0, 10)
	fmt.Printf("[DC] Sum(%v) = %d\n", arr, chapter4.Sum(arr))
	fmt.Printf("[DC] Count(%v) = %d\n", arr, chapter4.Count(arr))
	fmt.Printf("[DC] Max(%v) = %d\n", arr, chapter4.Max(arr))

	// Sort
	arr = helper.MakeRandomArray(10, -100, 100)
	fmt.Printf("[QuickSort] Unsorted array: %v\n", arr)
	fmt.Printf("[QuickSort] Sorted array: %v\n", chapter4.QuickSort(arr))
	fmt.Printf("[MergeSort] Sorted array: %v\n", chapter4.MergeSort(arr))
}

func examplesChapter5() {
	// Hash
	fmt.Printf("[Hash] SHA256(a) = %s\n", chapter5.CalculateHash("a"))
	fmt.Printf("[Hash] SHA256(b) = %s\n", chapter5.CalculateHash("b"))
}
