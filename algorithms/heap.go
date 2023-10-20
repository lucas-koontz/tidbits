package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


func buildHeapByInit(array []int) *MinHeap {
	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	*minHeap = array
	heap.Init(minHeap)
	fmt.Println("buildHeapByInit: ", *minHeap)
	return minHeap
}

func buildHeapByPush(array []int) *MinHeap {
	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	for _, num := range array {
		heap.Push(minHeap, num)
	}
	fmt.Println("buildHeapByPush: ", *minHeap)

	min := heap.Pop(minHeap)

	return minHeap
}


