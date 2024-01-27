package chapter9_test

import (
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter9"
)

func TestLongestCommonSubstring(t *testing.T) {
	tests := []struct {
		name, a, b, want string
	}{
		{"hish-vista", "vista", "hish", "is"},
		{"hish-fish", "fish", "hish", "ish"},
		{"blue-clues", "blue", "clues", "lue"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chapter9.LongestCommonSubstring(tt.a, tt.b); got != tt.want {
				t.Errorf("LongestCommonSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name, a, b string
		want       int
	}{
		{"fosh-fish", "fish", "fosh", 3},
		{"fosh-fort", "fort", "fosh", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chapter9.LongestCommonSubsequence(tt.a, tt.b); got != tt.want {
				t.Errorf("LongestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
