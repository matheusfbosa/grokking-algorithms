package chapter6

type Graph struct {
	Vertices map[string][]string
}

func NewGraph(vertices map[string][]string) *Graph {
	return &Graph{
		Vertices: vertices,
	}
}

// BreadthFirstSearch implements Breadth-First Search (BFS) algorithm to find salesperson in graph.
func (g *Graph) BreadthFirstSearch(target string) bool {
	searchQueue := make([]string, 0)
	searchQueue = append(searchQueue, g.Vertices[target]...)
	checked := make(map[string]bool)

	for len(searchQueue) > 0 {
		val := searchQueue[0]
		searchQueue = searchQueue[1:]

		if !checked[val] {
			if isSalesperson(val) {
				return true
			}

			searchQueue = append(searchQueue, g.Vertices[val]...)
			checked[val] = true
		}
	}

	return false
}

func isSalesperson(name string) bool {
	return string(name[len(name)-1]) == "m"
}
