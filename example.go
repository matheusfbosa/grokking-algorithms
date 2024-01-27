package main

import (
	"fmt"
	"math"

	"github.com/matheusfbosa/grokking-algorithms/chapter1"
	"github.com/matheusfbosa/grokking-algorithms/chapter2"
	"github.com/matheusfbosa/grokking-algorithms/chapter3"
	"github.com/matheusfbosa/grokking-algorithms/chapter4"
	"github.com/matheusfbosa/grokking-algorithms/chapter5"
	"github.com/matheusfbosa/grokking-algorithms/chapter6"
	"github.com/matheusfbosa/grokking-algorithms/chapter7"
	"github.com/matheusfbosa/grokking-algorithms/chapter8"
	"github.com/matheusfbosa/grokking-algorithms/helper"
)

func main() {
	//examplesChapter1()
	//examplesChapter2()
	//examplesChapter3()
	//examplesChapter4()
	//examplesChapter5()
	//examplesChapter6()
	//examplesChapter7()
	examplesChapter8()
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

func examplesChapter6() {
	// Graph
	vertices := map[string][]string{
		"you":    {"alice", "bob", "claire"},
		"bob":    {"anuj", "peggy"},
		"alice":  {"peggy"},
		"claire": {"thom", "jonny"},
		"anuj":   {},
		"peggy":  {},
		"thom":   {},
		"jonny":  {},
	}
	graph := chapter6.NewGraph(vertices)
	fmt.Printf("[BFS] Searching for salesperson in graph: %v\n", graph.Vertices)
	fmt.Printf("[BFS] Found: %t\n", graph.BreadthFirstSearch("you"))
}

func examplesChapter7() {
	// Dijkstra
	graph := map[string]map[string]int{
		"start": {"a": 6, "b": 2},
		"a":     {"end": 1},
		"b":     {"a": 3, "end": 5},
		"end":   {},
	}
	costs := map[string]float64{"a": 6, "b": 2, "end": math.Inf(1)}
	parents := map[string]string{
		"a":   "start",
		"b":   "start",
		"end": "",
	}

	dijkstra := chapter7.NewDijkstra(graph, costs, parents)
	dijkstra.Run()

	fmt.Println("[Dijkstra] Graph:", dijkstra.Graph)
	fmt.Println("[Dijkstra] Final costs:", dijkstra.Costs)
	fmt.Println("[Dijkstra] Final parents:", dijkstra.Parents)
}

func examplesChapter8() {
	// Greedy Algorithms
	statesToCover := map[string]bool{
		"mt": true, "wa": true, "or": true, "id": true, "nv": true, "ut": true, "ca": true, "az": true,
	}
	fmt.Println("[Greedy Algorithms] States to cover:", statesToCover)

	radioStations := map[string]map[string]bool{
		"kone":   {"id": true, "nv": true, "ut": true},
		"ktwo":   {"wa": true, "id": true, "mt": true},
		"kthree": {"or": true, "nv": true, "ca": true},
		"kfour":  {"nv": true, "ut": true},
		"kfive":  {"ca": true, "az": true},
	}

	finalRadioStations := chapter8.FindFinalRadioStations(statesToCover, radioStations)
	fmt.Println("[Greedy Algorithms] States to cover:", statesToCover)
	fmt.Println("[Greedy Algorithms] Radio stations:", radioStations)
	fmt.Println("[Greedy Algorithms] Final radio stations:", finalRadioStations)
}
