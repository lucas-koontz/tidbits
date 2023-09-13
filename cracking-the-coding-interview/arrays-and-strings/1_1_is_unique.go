package main

import (
	"fmt"
	"slices"
)

type Bytes []byte

func (b Bytes) Len() int           { return len(b) }
func (b Bytes) Less(i, j int) bool { return b[i] < b[j] }
func (b Bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func IsUnique(s string) bool {
	/*
		Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

		Approach 1: Hashmap
			Insert all characters into a hashmap. If a character already exists, return false. If all characters are unique, return true.
		Approach 2: Sort and compare
			Sort the string and compare adjacent characters. If any adjacent characters are the same, return false. If all characters are unique, return true.
	*/

	sAsBytes := []byte(s)

	// Approach 1
	/*
		hashMap := make(map[byte]bool)

		for _, char := range sAsBytes {
			if _, ok := hashMap[char]; ok {
				return false
			}
			hashMap[char] = true
		}
	*/

	// Approach 2
	// sort.Sort(Bytes(sAsBytes))
	slices.Sort(sAsBytes)
	for i := 1; i < len(sAsBytes); i++ {
		if sAsBytes[i] == sAsBytes[i-1] {
			return false
		}
	}

	return true
}

type IsUniqueTestCase struct {
	input    string
	expected bool
}

func TestIsUnique() {
	testCases := []IsUniqueTestCase{
		{"test", false},
		{"case", true},
		{"unique", false},
		{"abcdea", false},
		{"abcdef", true},
	}

	fmt.Println("\nIs Unique")
	for _, testCase := range testCases {
		fmt.Println(fmt.Sprintf("Input: '%s'\n\tResult:%t\n\tExpect:%t", testCase.input, IsUnique(testCase.input), testCase.expected))
	}
}
