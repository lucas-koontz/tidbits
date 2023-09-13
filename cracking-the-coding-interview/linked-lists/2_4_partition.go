package main

import "fmt"

func Partition(head *Node, x int) *Node {
	/*
		Write code to partition a linked list around a value x, such that all nodes less than x come before all nodes greater than or equal to x.
		If x is contained within the list, the values of x only need to be after the elements less than x.
		The partition element x can appear anywhere in the "right partition"; it does not need to appear between the left and right partitions.

		                               v
		3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1, x == 5

		rootLeft = 0 -> 3 -> 2 -> 1 -> 5
						          ^ cursorLeft
		rootRight = 0 -> 5 -> 8 -> 5 -> 10 (-> nil)
				                         ^ cursorRight



		curr = 3

		Time Complexity: O(n)
		Space Complexity: O(1)
	*/

	rootLeft := &Node{0, nil}
	rootRight := &Node{0, nil}

	curr := head

	cursorLeft, cursorRight := rootLeft, rootRight

	for curr != nil {
		if curr.Value < x {
			// add to left
			cursorLeft.Next = curr
			cursorLeft = curr
		} else {
			// add to the right
			cursorRight.Next = curr
			cursorRight = curr
		}

		curr = curr.Next
	}

	cursorRight.Next = nil
	cursorLeft.Next = rootRight.Next

	return rootLeft.Next
}

type TestCasePartition struct {
	input    *Node
	x        int
	expected *Node
}

func TestPartition() {
	testCases := []TestCasePartition{
		{
			input: &Node{3,
				&Node{5,
					&Node{8,
						&Node{5,
							&Node{10,
								&Node{2,
									&Node{1, nil}}}}}}},
			x: 5,
			expected: &Node{3,
				&Node{2,
					&Node{1,
						&Node{5,
							&Node{8,
								&Node{5,
									&Node{10, nil}}}}}}},
		},
	}

	fmt.Println("\nPartition")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v - 'x == %d'\n\tResult:%v\n\tExpect:%v",
				testCase.x,
				Display(testCase.input),
				Display(Partition(testCase.input, testCase.x)),
				Display(testCase.expected)))
	}
}
