package chapter4

func Count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + Count(arr[1:])
}
