package chapter4

func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	currentMax := Max(arr[1:])
	if currentMax > arr[0] {
		return currentMax
	}

	return arr[0]
}
