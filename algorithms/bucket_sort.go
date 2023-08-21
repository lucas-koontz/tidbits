package main

func sortArray(nums []int) []int {
	size := 5 * 2 * 10000

	bucket := make([]int, size+1)

	for _, num := range nums {
		bucket[num+50000]++
	}

	sortedArray := make([]int, len(nums))
	index := 0

	for i := 0; i < len(bucket); i++ {
		count := bucket[i]
		for j := 0; j < count; j++ {

			sortedArray[index] = i - 50000
			index++

		}
	}
	return sortedArray
}
