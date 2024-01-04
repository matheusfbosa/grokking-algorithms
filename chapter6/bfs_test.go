package chapter6

import "testing"

var vertices = map[string][]string{
	"you":    {"alice", "bob", "claire"},
	"bob":    {"anuj", "peggy"},
	"alice":  {"peggy"},
	"claire": {"thom", "jonny"},
	"anuj":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

func TestGraph_BreadthFirstSearch(t *testing.T) {
	tests := []struct {
		name   string
		target string
		want   bool
	}{
		{
			name:   "SalespersonFound",
			target: "you",
			want:   true,
		},
		{
			name:   "SalespersonNotFound",
			target: "alice",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				Vertices: vertices,
			}

			if got := g.BreadthFirstSearch(tt.target); got != tt.want {
				t.Errorf("Graph.BreadthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
