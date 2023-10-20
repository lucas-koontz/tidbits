type Deque []int

func (deque *Deque) append(x int) {
	*deque = append(*deque, x)
}

func (deque *Deque) popFirst() {
	*deque = (*deque)[1:]
}

func (deque *Deque) popLast() {
	*deque = (*deque)[:len(*deque)-1]
}

func (deque Deque) first() int {
	return deque[0]
}

func (deque Deque) last() int {
	return deque[len(deque)-1]
}

func (deque Deque) len() int {
	return len(deque)
}

func maxSlidingWindow(nums []int, k int) (windows []int) {
	n := len(nums)

	deque := &Deque{}

	for i := 0; i < n; i++ {
		for deque.len() > 0 && nums[i] >= nums[deque.last()] {
			deque.popLast()
		}

		deque.append(i)

		if (i - k + 1) > deque.first() {
			deque.popFirst()
		}

		if i+1 >= k {
			windows = append(windows, nums[deque.first()])
		}
	}

	return windows
}
