package main

import (
	"fmt"
	"sort"
)

func main() {
	// TestIsUnique()
	// TestPermutation()
	// TestURLify()
	//TestPalindromePermutation()
	// TestOneAway()
	//TestStringCompression()
	//TestRotateMatrix()
	//TesZeroMatrix()
	// TestStringRotation()

	str := "onlylowercasecharacters"
	strBytes := []byte(str)
	fmt.Println(string(strBytes))
	sort.Slice(strBytes, func(i, j int) bool {
		return str[i] < str[j]
	},
	)
	fmt.Println(string(strBytes))


	arr := []int{5, 6, 7, 3, 2, 1}
	fmt.Println(arr)
	
	
	sort.Sort(sort.IntSlice(arr))
	fmt.Println(arr)

	fmt.Println(sort.Sort(sort.IntSlice([]int(str))))

}

type array []int
	func (a array) Len() int {
		return len(a)
	}
	func (a array) Less(i, j int) bool {
		return a[i] < a[j]
	}
	func (a array) Swap(i, j int) {
		a[i], a[j] = a[j], a[i]
	}
