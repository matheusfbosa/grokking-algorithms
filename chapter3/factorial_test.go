package chapter3

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"0!", 0, 1},
		{"1!", 1, 1},
		{"2!", 2, 2},
		{"3!", 3, 6},
		{"4!", 4, 24},
		{"5!", 5, 120},
		{"10!", 10, 3628800},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.n); got != tt.want {
				t.Errorf("%d! = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
