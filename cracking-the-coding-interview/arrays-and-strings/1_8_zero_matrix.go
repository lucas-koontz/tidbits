package main

import "fmt"

func ZeroMatrix(matrix [][]int) [][]int {
	/*
		Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.

		// Approach: HashMap

			Time Complexity: O(nˆ2)
			Space Complexity: O(n+m)
	*/

	n := len(matrix)
	if n < 1 {
		return matrix
	}
	m := len(matrix[0])

	rows := make(map[int]bool)
	columns := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := rows[i]; ok {
				matrix[i][j] = 0
			} else if _, ok := columns[j]; ok {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

type TestCaseZeroMatrix struct {
	input    [][]int
	expected [][]int
}

func TesZeroMatrix() {
	testCases := []TestCaseZeroMatrix{
		{[][]int{}, [][]int{}},
		{[][]int{{1, 0}, {3, 4}}, [][]int{{0, 0}, {3, 0}}},
		{[][]int{{1, 0}, {0, 4}}, [][]int{{0, 0}, {0, 0}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{1, 2, 0}, {4, 5, 6}, {7, 8, 9}}, [][]int{{0, 0, 0}, {4, 5, 0}, {7, 8, 0}}},
	}

	fmt.Println("\nZero Matrix")
	for _, testCase := range testCases {
		fmt.Println(fmt.Sprintf("Input: '%d'\n\tResult:%d\n\tExpect:%d", testCase.input, ZeroMatrix(testCase.input), testCase.expected))
	}
}
