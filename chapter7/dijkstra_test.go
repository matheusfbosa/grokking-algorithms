package chapter7_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter7"
)

func TestDijkstra_Run(t *testing.T) {
	tests := []struct {
		name            string
		initialGraph    map[string]map[string]int
		initialCosts    map[string]float64
		initialParents  map[string]string
		expectedCosts   map[string]float64
		expectedParents map[string]string
	}{
		{
			name: "Example",
			initialGraph: map[string]map[string]int{
				"start": {"a": 6, "b": 2},
				"a":     {"end": 1},
				"b":     {"a": 3, "end": 5},
				"end":   {},
			},
			initialCosts: map[string]float64{"a": 6, "b": 2, "end": math.Inf(1)},
			initialParents: map[string]string{
				"a":   "start",
				"b":   "start",
				"end": "",
			},
			expectedCosts: map[string]float64{"a": 5, "b": 2, "end": 6},
			expectedParents: map[string]string{
				"a":   "b",
				"b":   "start",
				"end": "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := chapter7.NewDijkstra(tt.initialGraph, tt.initialCosts, tt.initialParents)
			d.Run()

			if !reflect.DeepEqual(d.Costs, tt.expectedCosts) {
				t.Errorf("Unexpected final costs. Got %v, expected %v", d.Costs, tt.expectedCosts)
			}

			if !reflect.DeepEqual(d.Parents, tt.expectedParents) {
				t.Errorf("Unexpected final parents. Got %v, expected %v", d.Parents, tt.expectedParents)
			}
		})
	}
}
