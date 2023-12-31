package chapter3_test

import (
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter3"
)

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
			if got := chapter3.Factorial(tt.n); got != tt.want {
				t.Errorf("%d! = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFactorialTailRecursion(t *testing.T) {
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
			if got := chapter3.FactorialTailRecursion(tt.n); got != tt.want {
				t.Errorf("%d! = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chapter3.Factorial(20)
	}
}

func BenchmarkFactorialTailRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chapter3.FactorialTailRecursion(20)
	}
}
