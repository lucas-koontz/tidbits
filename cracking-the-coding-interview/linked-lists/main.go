package main

import (
	b64 "encoding/base64"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func Display(head *Node) []int {
	n := 0

	hashMap := make(map[*Node]bool)

	curr := head
	for curr != nil {
		if _, ok := hashMap[curr]; ok {
			break
		}
		hashMap[curr] = true
		n++
		curr = curr.Next
	}

	output := make([]int, n)

	curr = head
	for i := 0; i < n; i++ {
		output[i] = curr.Value
		curr = curr.Next
	}

	return output
}

func main() {
	//TestRemoveDups()
	// TestReturnKthLast()
	// TestDeleteNodeMiddle()
	// TestPartition()
	// TestSumList()
	//TestSumListFollowUp()
	// TestPalindrome()
	//TestIntersection()
	//TestLoopDetection()
	data := "wazp_perforce@wildlifestudios.com:KhB$gRkg4x9PHFp884XiLq4Pnbaj@ah^2j"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
}
