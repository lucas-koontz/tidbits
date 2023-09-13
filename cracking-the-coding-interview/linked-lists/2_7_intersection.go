package main

import "fmt"

func Intersection(l1 *Node, l2 *Node) *Node {
	/*
		Given two (singly) linked lists, determine if the two lists intersect. Return the intersecting node.
		Note that the intersection is defined based on reference, not value. That is, if the kth node of the first linked list is the exact same node
		(by reference) as the jth node of the second linked list, then they are intersecting.

		Approach:
			Check last node reference. If same node, it intersects. Otherwise, it does not (return nil).
			If lists have the same length, run through both lists at the same time and find the equivalent node on both.
			Otherwise, check the difference in size i == n1 - n2, where n1 is the size of the longer list and n2 is the size of the smallest list.
				Return the (i - 1)th node.
	*/

	if l1 == nil || l2 == nil {
		return nil
	}

	n1, n2 := 0, 0

	var lastL1, lastL2 *Node
	currL1, currL2 := l1, l2

	// Find the last node
	for currL1 != nil {
		n1++
		lastL1 = currL1
		currL1 = currL1.Next
	}

	for currL2 != nil {
		n2++
		lastL2 = currL2
		currL2 = currL2.Next
	}

	if lastL1 != lastL2 { // Does not intersect
		return nil
	}

	if n1 == n2 { // Run through both at the same time
		currL1, currL2 := l1, l2
		for currL1 != currL2 {
			currL1 = currL1.Next
			currL2 = currL2.Next
		}
		return currL1
	} else {
		if n1 < n2 {
			n1, n2 = n2, n1
			l1, l2 = l2, l1
		}
		diff := n1 - n2
		count := 0

		currL2 := l2

		for count <= diff {
			currL2 = currL2.Next
			count++
		}

		return currL2
	}
}

type TestCaseIntersection struct {
	l1       *Node
	l2       *Node
	expected *Node
}

func TestIntersection() {
	f := &Node{6, nil}
	e := &Node{5, f}
	d := &Node{4, e}
	c := &Node{3, d}
	b := &Node{2, c}
	a := &Node{1, b}

	c2 := &Node{32, d}
	b2 := &Node{22, c2}

	c3 := &Node{33, d}
	b3 := &Node{23, c3}
	a3 := &Node{13, b3}

	testCases := []TestCaseIntersection{
		{
			l1:       a,
			l2:       b2,
			expected: d,
		}, {
			l1:       a,
			l2:       a3,
			expected: d,
		},
	}

	fmt.Println("\nIntersection")

	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v - %v \n\tResult:%v\n\tExpect:%v",
				Display(testCase.l1),
				Display(testCase.l2),
				Intersection(testCase.l1, testCase.l2),
				testCase.expected))
	}
}
