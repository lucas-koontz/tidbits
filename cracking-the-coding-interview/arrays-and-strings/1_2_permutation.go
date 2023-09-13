package main

import "fmt"

func Permutation(s1, s2 string) bool {
	/*
		Given two strings, write a method to decide if one is a permutation of the other.

		Approach: Hashmap
			Create a hashmap and store information about each string
	*/

	hashMap := make(map[int32]int)

	for _, char := range s1 {
		hashMap[char]++
	}

	for _, char := range s2 {
		hashMap[char]--
	}

	for _, value := range hashMap {
		if value != 0 {
			return false
		}
	}

	return true
}

type PermutationTestCase struct {
	string1  string
	string2  string
	expected bool
}

func TestPermutation() {
	testCases := []PermutationTestCase{
		{"test", "", false},
		{"", "test", false},
		{"tset", "test", true},
		{"aabbccdd", "abcdabcd", true},
		{"something", "another", false},
	}

	fmt.Println("\nPermutation")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: '%s' and '%s'\n\tResult:%t\n\tExpect:%t",
				testCase.string1, testCase.string2, Permutation(testCase.string1, testCase.string2), testCase.expected))
	}
}
