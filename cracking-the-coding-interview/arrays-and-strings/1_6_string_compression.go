package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type CharCount struct {
	char  byte
	count int
}

func StringCompression(s string) string {
	/*
		Implement a method to perform basic string compression using the counts of repeated characters.
		For example, the string 'aabcccccaaa' would become 'a2b1c5a3'.
		If the "compressed" string would not become smaller than the original string, your method should return the original string.
		You can assume the string has only uppercase and lowercase letters (a-z).

		Approach: Count how many characters repeat

		Time Complexity: O(n)
		Space Complexity: O(n) - if each character is unique in that space, it gives us O(2 * n) == O(n)
	*/

	if len(s) < 1 {
		return s
	}

	buffer := &bytes.Buffer{}

	charCount := CharCount{s[0], 1}

	for i := 1; i < len(s); i++ {
		if charCount.char == s[i] {
			charCount.count++
		} else {
			// Write to newS
			StringCompressionWrite(buffer, charCount)

			charCount = CharCount{s[i], 1}
		}
	}

	StringCompressionWrite(buffer, charCount)
	newS := buffer.String()

	if len(newS) > len(s) {
		return s
	} else {
		return newS
	}
}

func StringCompressionWrite(buffer *bytes.Buffer, charCount CharCount) {
	buffer.WriteString(string(charCount.char))
	buffer.WriteString(strconv.Itoa(charCount.count))
}

type TestCaseStringCompression struct {
	input    string
	expected string
}

func TestStringCompression() {
	testCases := []TestCaseStringCompression{
		{"aabcccccaaa", "a2b1c5a3"},
	}

	fmt.Println("\nStringCompression")
	for _, testCase := range testCases {
		fmt.Println(fmt.Sprintf("Input: '%s'\n\tResult:%s\n\tExpect:%s", testCase.input, StringCompression(testCase.input), testCase.expected))
	}
}
