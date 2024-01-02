package chapter4_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter4"
	"github.com/matheusfbosa/grokking-algorithms/helper"
)

var tests = []struct {
	name string
	arr  []int
	want []int
}{
	{
		name: "Empty",
		arr:  []int{},
		want: []int{},
	},
	{
		name: "Sorted",
		arr:  []int{1, 2, 3, 4, 5},
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "ReverseSorted",
		arr:  []int{5, 4, 3, 2, 1},
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "Duplicates",
		arr:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
		want: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
	},
	{
		name: "Random",
		arr:  []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 10},
		want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chapter4.QuickSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chapter4.MergeSort(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := helper.MakeRandomArray(1000, -1000, 1000)
		chapter4.QuickSort(arr)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := helper.MakeRandomArray(1000, -1000, 1000)
		chapter4.MergeSort(arr)
	}
}

func BenchmarkBuiltInSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := helper.MakeRandomArray(1000, -1000, 1000)
		sort.Ints(arr)
	}
}
