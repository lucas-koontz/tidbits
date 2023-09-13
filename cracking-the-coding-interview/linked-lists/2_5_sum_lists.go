package main

import "fmt"

func SumList(l1, l2 *Node) *Node {
	/*
		You have two numbers represented by a linked list, where each node contains a single digit.
		The digits are stored in reverse order, such that the 1's digit is at the head of the list.
		Write a function that adds the two numbers and returns the sum as a linked list.
		Example:
			Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
			Output: 2 -> 1 -> 9. That is, 912.

		FOLLOW UP
			Suppose the digits are stored in forward order. Repeat the above problem.
			Example:
				Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
				Output: 9 -> 1 -> 2. That is, 912.

	*/

	/*
		           v
		(7 -> 1 -> 6)
		           v
		(5 -> 9 -> 2)

		carry := 1

		root := 0 -> 2 -> 1 -> 9

		n1, n2 = 7, 5
		sum := 7 + 5 + 0 == 12
		carry := 12 / 10 == 1

		node := { 2, nil }


		n1, n2 = 1, 9
		sum := 1 + 9 + 1 == 11
		carry := 11 / 10 == 1

		node := { 1, nil }


		n1, n2 = 6, 2
		sum := 6 + 2 + 1 == 9
		carry := 9 / 10 == 0

		node := { 9, nil }

	*/

	cursorL1, cursorL2 := l1, l2
	carry := 0

	root := &Node{}
	curr := root

	for cursorL1 != nil || cursorL2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if cursorL1 != nil {
			n1 = cursorL1.Value
			cursorL1 = cursorL1.Next
		}
		if cursorL2 != nil {
			n2 = cursorL2.Value
			cursorL2 = cursorL2.Next
		}

		sum := n1 + n2 + carry
		carry = sum / 10

		node := &Node{
			Value: sum % 10,
			Next:  nil,
		}

		curr.Next = node
		curr = node
	}

	return root.Next
}

type TestCaseSumList struct {
	l1       *Node
	l2       *Node
	expected *Node
}

func TestSumList() {
	testCases := []TestCaseSumList{
		{
			l1: &Node{7,
				&Node{1,
					&Node{6, nil}}},
			l2: &Node{5,
				&Node{9,
					&Node{2, nil}}},
			expected: &Node{2,
				&Node{1,
					&Node{9, nil}}},
		},
		{
			l1: &Node{7,
				&Node{1, nil}},
			l2: &Node{5,
				&Node{9,
					&Node{2, nil}}},
			expected: &Node{2,
				&Node{1,
					&Node{3, nil}}},
		},
	}

	fmt.Println("\nSum List")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v and %v\n\tResult:%v\n\tExpect:%v",
				Display(testCase.l1),
				Display(testCase.l2),
				Display(SumList(testCase.l1, testCase.l2)),
				Display(testCase.expected)))
	}
}
