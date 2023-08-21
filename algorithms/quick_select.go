package main

import "fmt"

func main() {
	numbers := []int{1, 7, 2, 8, 5, 9, 3, 6, 4}

	k := 2
	QuickSelect(numbers, 0, len(numbers)-1, k)
	fmt.Println("kth:", k, "number:", numbers[k-1], "numbers[0:k]:", numbers[:k])

	k = 5
	QuickSelect(numbers, 0, len(numbers)-1, k)
	fmt.Println("kth:", k, "number:", numbers[k-1], "numbers[0:k]:", numbers[:k])

	k = 7
	numbers = []int{-5, -9, 67, -10, 4, -57, 47, 13, -67, -26, -57, 63, 38, -68, 62, -45, -37, -100, 95, 49, -91, -53}
	QuickSelect(numbers, 0, len(numbers)-1, k)
	fmt.Println("kth:", k, "number:", numbers[k-1], "numbers[0:k]:", numbers[:k])
}

func QuickSelect(numbers []int, left, right, k int) {
	if left < right {
		pivot := partition(numbers, left, right)
		if k == pivot {
			return
		} else if k < pivot {
			QuickSelect(numbers, left, pivot-1, k)
		} else {
			QuickSelect(numbers, pivot+1, right, k)
		}
	}
}

func partition(numbers []int, left, right int) int {
	pivot := numbers[right]

	i := left - 1
	for j := left; j <= right-1; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	i++
	numbers[i], numbers[right] = numbers[right], numbers[i]
	return i
}
