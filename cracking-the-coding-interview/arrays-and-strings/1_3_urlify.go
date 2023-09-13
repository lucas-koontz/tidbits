package main

import "fmt"

func URLify(str string, length int) string {
	/*
		Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are given the 'true' length of the string.

		Approach: Iterate backwards
			Iterate backwards through the string. If the character is a space, replace it with '%20'. If the character is not a space, move it to the end of the string.
	*/

	// Convert string to byte slice
	strAsBytes := []byte(str)

	index := len(str)

	// Iterate backwards through the string
	for i := length - 1; i >= 0; i-- {
		if strAsBytes[i] == ' ' {
			// Replace space with '%20'
			strAsBytes[index-1] = '0'
			strAsBytes[index-2] = '2'
			strAsBytes[index-3] = '%'
			index -= 3
		} else {
			// Move character to end of string
			strAsBytes[index-1] = strAsBytes[i]
			index--
		}
	}

	return string(strAsBytes)
}

type URLifyTestCase struct {
	str      string
	length   int
	expected string
}

func TestURLify() {
	testCases := []URLifyTestCase{
		{"Mr John Smith    ", 13, "Mr%20John%20Smith"},
	}

	fmt.Println("\nPermutation")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: '%s' and '%d'\n\tResult:%s\n\tExpect:%s",
				testCase.str, testCase.length, URLify(testCase.str, testCase.length), testCase.expected))
	}
}
