package main

import "fmt"

func Palindrome(head *Node) bool {
	/*
			Check if linked list is palindrome.

			Approach: Recursion or Copy lists iterative

			Time Complexity: O(n)
			Space Complexity: O(n) - Recursion
		                      v
			head -> 2 -> 3 -> 1 -> 3 -> 2 -> nil
					          ^
			cmpHead -> head
	*/

	if head == nil {
		return true
	}
	/*
		_, isPalindrome := IsPalindrome(head, head)

		return isPalindrome

	*/
	var cloneHead *Node

	curr := head

	for curr != nil {
		node := &Node{
			curr.Value,
			cloneHead,
		}
		cloneHead = node

		curr = curr.Next
	}

	curr1, curr2 := head, cloneHead
	for curr1 != nil { // && curr != nil
		if curr1.Value != curr2.Value {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
	return true
}

func IsPalindrome(node *Node, head *Node) (*Node, bool) {
	var cmpNode *Node
	shouldContinue := true

	if node.Next == nil {
		cmpNode = head
	} else {
		cmpNode, shouldContinue = IsPalindrome(node.Next, head)
	}

	if !shouldContinue {
		return nil, false
	}

	if node.Value == cmpNode.Value {
		return cmpNode.Next, true
	} else {
		return nil, false
	}
}

type TestCaseSumPalindrome struct {
	input    *Node
	expected bool
}

func TestPalindrome() {
	testCases := []TestCaseSumPalindrome{
		{
			input: &Node{2,
				&Node{3,
					&Node{7,
						&Node{3,
							&Node{2, nil}}}}},
			expected: true,
		},
		{
			input: &Node{2,
				&Node{3,
					&Node{7,
						&Node{7,
							&Node{3,
								&Node{2, nil}}}}}},
			expected: true,
		},
		{
			input: &Node{1,
				&Node{3,
					&Node{7,
						&Node{3,
							&Node{2, nil}}}}},
			expected: false,
		},
		{
			input: &Node{2,
				&Node{3,
					&Node{3,
						&Node{7,
							&Node{2,
								&Node{2, nil}}}}}},
			expected: false,
		},
	}

	fmt.Println("\nPalindrome")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v \n\tResult:%v\n\tExpect:%v",
				Display(testCase.input),
				Palindrome(testCase.input),
				testCase.expected))
	}
}
