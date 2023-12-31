package chapter4_test

import (
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter4"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "NonEmpty",
			arr:  []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "Empty",
			arr:  []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chapter4.Count(tt.arr); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
