package main

import "fmt"

func LoopDetection(head *Node) *Node {
	/*
		Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.
		DEFINITION
			Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node,
			so as to make a loop in the linked list.
		EXAMPLE
			Input: A -> B -> C -> D -> E -> C [the same C as earlier]
			Output: C


		Approach: Runner Technique
			Time Complexity: O(n)
			Space Complexity: O(1)
	*/

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil { // Not a cycle
		return nil
	}

	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	// Both points to the same node
	return slow
}

type TestCaseLoopDetection struct {
	input    *Node
	expected *Node
}

func TestLoopDetection() {
	f := &Node{6, nil}
	e := &Node{5, f}
	d := &Node{4, e}
	c := &Node{3, d}
	b := &Node{2, c}
	a := &Node{1, b}

	f.Next = c

	f2 := &Node{6, nil}
	e2 := &Node{5, f2}
	d2 := &Node{4, e2}
	c2 := &Node{3, d2}
	b2 := &Node{2, c2}
	a2 := &Node{1, b2}

	f2.Next = d2

	testCases := []TestCaseLoopDetection{
		{
			input:    a,
			expected: c,
		},
		{
			input:    a2,
			expected: d2,
		},
	}

	fmt.Println("\nLoop Detection1")

	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v \n\tResult:%v\n\tExpect:%v",
				Display(testCase.input),
				Display(LoopDetection(testCase.input)),
				Display(testCase.expected)),
		)
	}
}
