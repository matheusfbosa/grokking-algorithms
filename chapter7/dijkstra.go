package chapter7

import "math"

type Dijkstra struct {
	Graph   map[string]map[string]int
	Costs   map[string]float64
	Parents map[string]string
}

func NewDijkstra(
	graph map[string]map[string]int, costs map[string]float64, parents map[string]string,
) *Dijkstra {
	return &Dijkstra{
		Graph:   graph,
		Costs:   costs,
		Parents: parents,
	}
}

func (d *Dijkstra) Run() {
	var processed []string
	node := d.findLowestCostNode(d.Costs, processed)

	for node != "" {
		cost := d.Costs[node]
		neighbors := d.Graph[node]

		for n, weight := range neighbors {
			newCost := cost + float64(weight)
			if d.Costs[n] > newCost {
				d.Costs[n] = newCost
				d.Parents[n] = node
			}
		}

		processed = append(processed, node)
		node = d.findLowestCostNode(d.Costs, processed)
	}
}

func (d *Dijkstra) findLowestCostNode(costs map[string]float64, processed []string) string {
	lowestCost := math.Inf(1)
	lowestCostNode := ""

	for node, cost := range costs {
		if cost < lowestCost && !contains(processed, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}

	return false
}
