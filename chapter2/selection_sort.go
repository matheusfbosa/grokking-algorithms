package chapter2

func SelectionSort(arr []int) []int {
	sorted := make([]int, 0, len(arr))

	for len(arr) > 0 {
		lowestIdx := findLowestIdx(arr)
		sorted = append(sorted, arr[lowestIdx])
		arr = append(arr[:lowestIdx], arr[lowestIdx+1:]...)
	}

	return sorted
}

func findLowestIdx(arr []int) int {
	lowest := arr[0]
	lowestIdx := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < lowest {
			lowest = arr[i]
			lowestIdx = i
		}
	}

	return lowestIdx
}
