package chapter9

import "fmt"

func LongestCommonSubstring(a, b string) string {
	longestLength := 0
	lastSubIndex := 0
	table := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				table[i][j] = table[i-1][j-1] + 1

				if table[i][j] > longestLength {
					longestLength = table[i][j]
					lastSubIndex = i
				}
			} else {
				table[i][j] = 0
			}
		}
	}

	fmt.Printf("[Dynamic Programming] Table: %v\n", table)
	return a[lastSubIndex-longestLength : lastSubIndex]
}

func LongestCommonSubsequence(a, b string) int {
	table := createMatrix(len(a)+1, len(b)+1)

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				table[i][j] = max(table[i-1][j], table[i][j-1])
			}
		}
	}

	fmt.Printf("[Dynamic Programming] Table: %v\n", table)
	return table[len(a)][len(b)]
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)

	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	return matrix
}
