package chapter1_test

import (
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter1"
)

func TestNode_Insert(t *testing.T) {
	t.Run("EmptyTree", func(t *testing.T) {
		var root *chapter1.Node
		val := 5
		root = root.Insert(val)
		if root.Val != val {
			t.Errorf("Expected root value %d, got %d", val, root.Val)
		}
	})

	t.Run("InsertSmallerToLeft", func(t *testing.T) {
		var root *chapter1.Node
		val1 := 5
		root = root.Insert(val1)

		val2 := 3
		root.Insert(val2)
		if root.Left.Val != val2 {
			t.Errorf("Expected left child value %d, got %d", val2, root.Left.Val)
		}
	})

	t.Run("InsertLargerToRight", func(t *testing.T) {
		var root *chapter1.Node
		val1 := 5
		root = root.Insert(val1)

		val2 := 8
		root.Insert(val2)
		if root.Right.Val != val2 {
			t.Errorf("Expected right child value %d, got %d", val2, root.Right.Val)
		}
	})
}

func TestNode_Search(t *testing.T) {
	var root *chapter1.Node
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, val := range values {
		root = root.Insert(val)
	}

	t.Run("ExistingValue", func(t *testing.T) {
		val := 6
		result := root.Search(val)
		if result == nil || result.Val != val {
			t.Errorf("Expected to find value %d in the tree, but got %v", val, result)
		}
	})

	t.Run("NonExistingValue", func(t *testing.T) {
		val := 11
		result := root.Search(val)
		if result != nil {
			t.Errorf("Expected nil result for value %d, but got %v", val, result)
		}
	})
}

func TestNode_InOrderTraversal(t *testing.T) {
	var root *chapter1.Node
	keys := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, key := range keys {
		root = root.Insert(key)
	}

	expectedOrder := []int{1, 3, 4, 6, 7, 8, 10, 13, 14}
	resultOrder := inOrderTraversalToArray(root)
	for i, val := range resultOrder {
		if val != expectedOrder[i] {
			t.Errorf("In-order traversal mismatch at index %d. Expected %d, got %d", i, expectedOrder[i], val)
		}
	}
}

func inOrderTraversalToArray(root *chapter1.Node) []int {
	var result []int
	if root != nil {
		result = append(result, inOrderTraversalToArray(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inOrderTraversalToArray(root.Right)...)
	}
	return result
}
