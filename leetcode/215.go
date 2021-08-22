package leetcode

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int { return len(h)}
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	if h.Len() < 1 {
		panic("No element")
	}
	lastElement := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return lastElement
}

func findKthLargest(nums []int, k int) int {
	maxHeap := IntHeap(nums[:k])
	heap.Init(&maxHeap)

	for _, e := range nums[k:] {
		heap.Push(&maxHeap, e)
		heap.Pop(&maxHeap)
	}

	return maxHeap.Pop().(int)
}