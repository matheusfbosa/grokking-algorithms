package chapter2

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"Unsorted", []int{64, 25, 12, 22, 11}, []int{11, 12, 22, 25, 64}},
		{"AlreadySorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"Unique", []int{1, 1, 1}, []int{1, 1, 1}},
		{"Empty", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SelectionSort(tt.input)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("Expected %v, got %v", tt.output, result)
			}
		})
	}
}
