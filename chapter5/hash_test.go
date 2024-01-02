package chapter5_test

import (
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter5"
)

func TestCalculateHash(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Hello",
			s:    "Hello",
			want: "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chapter5.CalculateHash(tt.s); got != tt.want {
				t.Errorf("CalculateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
