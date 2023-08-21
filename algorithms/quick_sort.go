package main

import (
	"math/rand"
)

func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)

	return nums
}

func QuickSort(nums []int, low, high int) {
	if low >= high {
		return
	}
	pivot := Partition(nums, low, high)
	QuickSort(nums, low, pivot-1)
	QuickSort(nums, pivot+1, high)
}

func Partition(nums []int, low, high int) int {
	r := rand.Intn(high-low) + low
	nums[r], nums[high] = nums[high], nums[r]

	pivot := nums[high]

	index := low - 1

	for i := low; i < high; i++ {
		if nums[i] < pivot {
			index++
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	index++
	nums[high], nums[index] = nums[index], nums[high]

	return index
}
