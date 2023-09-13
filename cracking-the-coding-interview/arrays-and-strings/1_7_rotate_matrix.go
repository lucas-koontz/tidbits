package main

import "fmt"

type Coordinate struct {
	i int
	j int
}

func RotateMatrix(matrix [][]int) [][]int {
	/*
		Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees.
		Can you do this in place?

		Approach 1: New Matrix
	*/

	n := len(matrix)

	if n < 2 {
		return matrix
	}

	x := n / 2
	y := x

	if n%2 == 1 {
		x++
	}

	// Approach : In Place Rotation
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {

			// n-(n-j-1)-1 = n - n + j + 1 - 1
			// n-(n-i-1) -1 = n - n + i +1 -1

			// i, j -> n-j-1, i -> n-i-1 , n-j-1 -> n-(n-j-1)-1 == j , n-i-1 -> (back to beginning) n-(n-i-1) -1 == i  , j

			pos1 := Coordinate{i, j}
			pos2 := Coordinate{n - j - 1, i}
			pos3 := Coordinate{n - i - 1, n - j - 1}
			pos4 := Coordinate{j, n - i - 1}

			temp := matrix[pos1.i][pos1.j]
			matrix[pos1.i][pos1.j] = matrix[pos4.i][pos4.j]
			matrix[pos4.i][pos4.j] = matrix[pos3.i][pos3.j]
			matrix[pos3.i][pos3.j] = matrix[pos2.i][pos2.j]
			matrix[pos2.i][pos2.j] = temp
		}
	}

	// Approach 1 : New Matrix
	/*
		newMatrix := make([][]int, n)
		for i := range newMatrix {
			newMatrix[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {

				newI := n - j - 1
				newJ := i

				newMatrix[newI][newJ] = matrix[i][j]

			}
		}
	*/

	return matrix
}

type TestCaseRotateMatrix struct {
	input    [][]int
	expected [][]int
}

func TestRotateMatrix() {
	testCases := []TestCaseRotateMatrix{
		{[][]int{}, [][]int{}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{2, 4}, {1, 3}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}, [][]int{{4, 8, 12, 16}, {3, 7, 11, 15}, {2, 6, 10, 14}, {1, 5, 9, 13}}},
	}

	fmt.Println("\nRotate Matrix")
	for _, testCase := range testCases {
		fmt.Println(fmt.Sprintf("Input: '%d'\n\tResult:%d\n\tExpect:%d", testCase.input, RotateMatrix(testCase.input), testCase.expected))
	}
}
