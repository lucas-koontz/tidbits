package main

import "fmt"

func RemoveDups(head *Node) *Node {
	/*
		Write code to remove duplicates from an unsorted linked list.
		FOLLOW UP
		How would you solve this problem if a temporary buffer is not allowed?

		Approach: HashMap
			Time Complexity: O(n)
			Space Complexity: O(m), where m is how many distinct elements exists in the list

		Approach 2: No buffer
			Time Complexity: O(nˆ2)
			Space Complexity: O(1)

	*/

	// Approach 1

	/*
		hashMap := make(map[int]bool)

		curr := head
		var prev *Node // dummy

		for curr != nil {
			if _, ok := hashMap[curr.Value]; ok {
				prev.Next = curr.Next
			}
			hashMap[curr.Value] = true
			prev = curr
			curr = curr.Next
		}

		return head
	*/

	// Approach 2
	curr := head

	for curr != nil {
		runner := curr.Next
		runnerPrev := curr

		for runner != nil {
			if runner.Value == curr.Value {
				runnerPrev.Next = runner.Next
			}
			runnerPrev = runner
			runner = runner.Next
		}
		curr = curr.Next
	}

	return head
}

type TestCaseRemoveDups struct {
	input    *Node
	expected *Node
}

func TestRemoveDups() {
	testCases := []TestCaseRemoveDups{
		{
			input:    &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			expected: &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
		},
		{
			input: &Node{1,
				&Node{1,
					&Node{3,
						&Node{4,
							&Node{5,
								&Node{1,
									&Node{2, nil}}}}}}},
			expected: &Node{1,
				&Node{3,
					&Node{4,
						&Node{5,
							&Node{2, nil}}}}},
		},
	}

	fmt.Println("\nRemove Dups")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v\n\tResult:%v\n\tExpect:%v",
				Display(testCase.input),
				Display(RemoveDups(testCase.input)),
				Display(testCase.expected)))
	}
}
