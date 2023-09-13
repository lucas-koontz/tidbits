package main

import "fmt"

func DeleteMiddleNode(node *Node) {
	/*
			Implement an algorithm to delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list,
		given only access to that node.


			a -> b -> c -> d -> e -> f

			we are given node 'c' - 'b.next' points to the address where c. is
	*/

	node.Value = node.Next.Value
	node.Next = node.Next.Next
}

type TestCaseDeleteNodeMiddle struct {
	input    *Node
	expected *Node
}

func TestDeleteNodeMiddle() {
	f := &Node{6, nil}
	e := &Node{5, f}
	d := &Node{4, e}
	c := &Node{3, d}
	b := &Node{2, c}
	a := &Node{1, b}

	testCases := []TestCaseDeleteNodeMiddle{
		{
			input:    c,
			expected: a,
		},
	}

	fmt.Println("\nDelete Node Middle")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v",
				Display(testCase.input)),
		)

		DeleteMiddleNode(testCase.input)

		fmt.Println(
			fmt.Sprintf("\n\tExpect:%v",
				Display(testCase.expected)),
		)
	}
}
