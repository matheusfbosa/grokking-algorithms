package chapter1

func BinarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		val := arr[mid]
		if val == target {
			return mid
		}

		if val > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
