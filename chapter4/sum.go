package chapter4

func Sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + Sum(arr[1:])
}
