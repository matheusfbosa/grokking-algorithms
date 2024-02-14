package chapter10

import "math"

func euclideanDistance(point1, point2 []float64) float64 {
	var sum float64
	for i := range point1 {
		sum += math.Pow(point1[i]-point2[i], 2)
	}
	return math.Sqrt(sum)
}

func KNN(xTrain [][]float64, yTrain []string, xTest []float64, k int) string {
	distances := make([]struct {
		dist  float64
		label string
	}, len(xTrain))

	for i, xTrainPoint := range xTrain {
		distances[i] = struct {
			dist  float64
			label string
		}{
			dist:  euclideanDistance(xTest, xTrainPoint),
			label: yTrain[i],
		}
	}

	for i := 0; i < k; i++ {
		minIndex := i
		for j := i + 1; j < len(distances); j++ {
			if distances[j].dist < distances[minIndex].dist {
				minIndex = j
			}
		}
		distances[i], distances[minIndex] = distances[minIndex], distances[i]
	}

	classCount := make(map[string]int)
	for _, neighbor := range distances[:k] {
		classCount[neighbor.label]++
	}

	var predictedClass string
	maxCount := 0
	for label, count := range classCount {
		if count > maxCount {
			maxCount = count
			predictedClass = label
		}
	}

	return predictedClass
}
