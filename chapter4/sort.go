package chapter4

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var less, greater []int

	for _, x := range arr[1:] {
		if x <= pivot {
			less = append(less, x)
		} else {
			greater = append(greater, x)
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftHalf := MergeSort(arr[:mid])
	rightHalf := MergeSort(arr[mid:])

	return merge(leftHalf, rightHalf)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
