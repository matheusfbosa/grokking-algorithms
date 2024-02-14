package chapter10_test

import (
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter10"
)

func TestKNN(t *testing.T) {
	tests := []struct {
		name   string
		xTrain [][]float64
		yTrain []string
		xTest  []float64
		k      int
		want   string
	}{
		{
			name:   "k3",
			xTrain: [][]float64{{1, 2}, {2, 3}, {3, 1}, {4, 2}, {5, 4}, {6, 5}, {7, 6}, {8, 5}},
			yTrain: []string{"A", "A", "A", "A", "B", "B", "B", "B"},
			xTest:  []float64{4, 4},
			k:      3,
			want:   "A",
		},
		{
			name:   "k5",
			xTrain: [][]float64{{1, 2}, {2, 3}, {3, 1}, {4, 2}, {5, 4}, {6, 5}, {7, 6}, {8, 5}},
			yTrain: []string{"A", "A", "A", "A", "B", "B", "B", "B"},
			xTest:  []float64{6, 6},
			k:      5,
			want:   "B",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chapter10.KNN(tt.xTrain, tt.yTrain, tt.xTest, tt.k); got != tt.want {
				t.Errorf("KNN() = %v, want %v", got, tt.want)
			}
		})
	}
}
