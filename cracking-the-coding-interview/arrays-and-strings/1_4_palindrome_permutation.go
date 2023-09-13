package main

import "fmt"

func PalindromePermutation(str string) bool {
	/*
		Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.

		Approach: Hashmap

		Input:
			Assuming only lowercase characters from alphabet.

		Two cases:
			1. Even number of characters
				Same amount of all characters
			2. Odd number of characters
				Same amount of all characters but one

		"aaa"
	*/

	strBytes := []byte(str)

	hashMap := make(map[byte]int)

	for _, char := range strBytes {
		hashMap[char]++
	}

	amountMap := make(map[int]bool)

	for _, value := range hashMap {
		amountMap[value] = true
	}

	if len(amountMap) > 2 {
		return false
	} else if len(amountMap) == 2 && len(str)%2 == 0 {
		return false
	} else {
		return true
	}
}

type PalindromePermutationTestCase struct {
	input    string
	expected bool
}

func TestPalindromePermutation() {
	testCases := []PalindromePermutationTestCase{
		{"test", false},
		{"tactcoa", true},
		{"tacocat", true},
		{"baa", true},
		{"abcdeabcde", true},
	}

	fmt.Println("\nPalindrome Permutation")
	for _, testCase := range testCases {
		fmt.Println(fmt.Sprintf("Input: '%s'\n\tResult:%t\n\tExpect:%t", testCase.input, PalindromePermutation(testCase.input), testCase.expected))
	}
}
