package main

import "fmt"

func sortArray(nums []int) []int {
	return MergeSort(nums)
}

func MergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	left := MergeSort(nums[:len(nums)/2])
	right := MergeSort(nums[len(nums)/2:])
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	sorted := make([]int, len(left)+len(right))
	fmt.Println(len(sorted))

	i := 0
	j := 0

	index := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted[index] = left[i]
			i++
		} else {
			sorted[index] = right[j]
			j++
		}
		index++
	}

	for ; i < len(left); i++ {
		sorted[index] = left[i]
		index++
	}

	for ; j < len(right); j++ {
		sorted[index] = right[j]
		index++
	}

	return sorted
}
