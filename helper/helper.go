package helper

import "math/rand"

func MakeRandomArray(size, min, max int) []int {
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}

	return arr
}

func SetIntersection(setA, setB map[string]bool) map[string]bool {
	result := make(map[string]bool)

	for key := range setA {
		if setB[key] {
			result[key] = true
		}
	}

	return result
}
