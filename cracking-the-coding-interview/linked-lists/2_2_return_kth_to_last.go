package main

import "fmt"

func ReturnKthLast(head *Node, k int) *Node {
	/*
		Implement an algorithm to find the kth to last element of a singly linked list.

		Approach: Find number of elements
			Assuming k is 0-indexed

		Approach: Two Pointers
	*/

	if head == nil {
		return nil
	}

	fast := head

	for ; k >= 0; k-- {
		fast = fast.Next
	}

	slow := head

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow

	//n := 0
	//curr := head
	//
	//for curr != nil {
	//	n++
	//	curr = curr.Next
	//}
	//
	//if k > n-1 || k < 0 {
	//	return nil
	//}
	//
	//target := n - k
	//count := 0
	//
	//curr = head
	//for curr != nil {
	//	count++
	//	if count == target {
	//		return curr
	//	}
	//	curr = curr.Next
	//}

	return nil
}

type TestCaseReturnKthLast struct {
	input    *Node
	k        int
	expected *Node
}

func TestReturnKthLast() {
	testCases := []TestCaseReturnKthLast{
		{
			input:    &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:        0,
			expected: &Node{5, nil},
		},
		{
			input:    &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:        1,
			expected: &Node{4, &Node{5, nil}},
		},
		{
			input:    &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, nil}}}}},
			k:        2,
			expected: &Node{3, &Node{4, &Node{5, nil}}},
		},
	}

	fmt.Println("\nReturn Kth to Last")
	for _, testCase := range testCases {
		fmt.Println(
			fmt.Sprintf("Input: %v - k: '%d'\n\tResult:%v\n\tExpect:%v",
				Display(testCase.input),
				testCase.k,
				Display(ReturnKthLast(testCase.input, testCase.k)),
				Display(testCase.expected)))
	}
}
