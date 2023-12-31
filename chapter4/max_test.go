package chapter4

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "Positive",
			arr:  []int{1, 9, 2, 8, 3},
			want: 9,
		},
		{
			name: "Empty",
			arr:  []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.arr); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
