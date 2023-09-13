package main

import (
	"fmt"
	"strings"
)

func StringRotation(s1, s2 string) bool {
	/*
		Assume you have a method isSubstring which checks if one word is a substring of another.
		Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring.


		Approach: Duplicate s1 on itself

	*/

	if len(s1) != len(s2) || len(s1) < 1 {
		return false
	}

	s := s1 + s1 // Concatenating only once, no buffer required
	return isSubstring(s, s2)
}

func isSubstring(s, substr string) bool {
	return strings.Contains(s, substr)
}

type TestCaseStringRotation struct {
	string1  string
	string2  string
	expected bool
}

func TestStringRotation() {
	testCases := []TestCaseStringRotation{
		{"waterbottle", "bottlewater", true},
		{"waterbottle", "waterbottle", true},
		{"waterbottle", "waterbottl", false},
	}

	fmt.Println("\nString Rotation")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: '%s' and '%s'\n\tResult:%t\n\tExpect:%t",
				testCase.string1, testCase.string2, StringRotation(testCase.string1, testCase.string2), testCase.expected))
	}
}
