package main

import "fmt"

func SumListFollowUp(l1, l2 *Node) *Node {
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

		Assuming at least 1 node on each list

		Pad the smallest list with 0 nodes
	*/

	n1, n2 := 0, 0

	cursor1, cursor2 := l1, l2

	for cursor1 != nil || cursor2 != nil {
		if cursor1 != nil {
			n1++
			cursor1 = cursor1.Next
		}
		if cursor2 != nil {
			n2++
			cursor2 = cursor2.Next
		}
	}

	if n1 < n2 { // could do |abs|
		l1, l2 = l2, l1
	}

	cursor2 = l2
	for i := 0; i < n1-n2; i++ {
		node := &Node{
			0,
			cursor2,
		}
		cursor2 = node
	}
	l2 = cursor2

	head, carry := SumRecursion(l1, l2)
	if carry > 0 {
		node := &Node{
			Value: carry,
			Next:  head,
		}
		head = node
	}

	return head
}

func SumRecursion(cursorL1, cursorL2 *Node) (*Node, int) {
	var next *Node
	var carry int

	if cursorL1.Next == nil && cursorL2.Next == nil { // Got to the end L1 and L2
		sum := cursorL1.Value + cursorL2.Value

		return &Node{
			Value: sum % 10,
			Next:  nil,
		}, sum / 10
	} else { // Both still going - Both are the same size
		next, carry = SumRecursion(cursorL1.Next, cursorL2.Next)
	}

	n1 := cursorL1.Value
	n2 := cursorL2.Value

	sum := n1 + n2 + carry
	carry = sum / 10

	return &Node{
		Value: sum % 10,
		Next:  next,
	}, carry
}

type TestCaseSumListFollowUp struct {
	l1       *Node
	l2       *Node
	expected *Node
}

func TestSumListFollowUp() {
	testCases := []TestCaseSumListFollowUp{
		//{
		//	l1: &Node{6,
		//		&Node{1,
		//			&Node{7, nil}}},
		//	l2: &Node{2,
		//		&Node{9,
		//			&Node{5, nil}}},
		//	expected: &Node{9,
		//		&Node{1,
		//			&Node{2, nil}}},
		//},
		{
			l1: &Node{6,
				&Node{1,
					&Node{7,
						&Node{4, nil}}}},
			l2: &Node{2,
				&Node{9,
					&Node{5, nil}}},
			expected: &Node{6,
				&Node{4,
					&Node{6,
						&Node{}}}},
		},
	}

	fmt.Println("\nSum List Follow Up")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v and %v\n\tResult:%v\n\tExpect:%v",
				Display(testCase.l1),
				Display(testCase.l2),
				Display(SumListFollowUp(testCase.l1, testCase.l2)),
				Display(testCase.expected)))
	}
}
