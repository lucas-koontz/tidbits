package main

import "fmt"

func OneAway(s1, s2 string) bool {
	/*
		There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character.
		Given two strings, write a function to check if they are one edit (or zero edits) away.


		Approach: Comparing
			Find the smallest.
			Go through the strings and find different.
				If same size, assume update.
				Otherwise, 'add' string.

			At the end, comparing remaining character from biggest string
	*/

	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		s1, s2 = s2, s1
		len1, len2 = len2, len1
	}

	operations := 0

	i, j := 0, 0
	for i < len1 && j < len2 {
		// Same size update. Otherwise, add.
		if s1[i] != s2[j] {
			if len1 == len2 {
				i++
				j++
			} else {
				j++
			}

			operations++
		} else {
			i++
			j++
		}
	}

	operations += len2 - j

	if operations > 1 {
		return false
	} else {
		return true
	}
}

type TestCaseOneAway struct {
	string1  string
	string2  string
	expected bool
}

func TestOneAway() {
	testCases := []TestCaseOneAway{
		{"pale", "ple", true},
		{"pale", "ale", true},
		{"pale", "pal", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"palesy", "pale", false},
		{"palesy", "pelican", false},
	}

	fmt.Println("\nOne Away")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: '%s' and '%s'\n\tResult:%t\n\tExpect:%t",
				testCase.string1, testCase.string2, OneAway(testCase.string1, testCase.string2), testCase.expected))
	}
}
