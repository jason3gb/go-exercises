package leetcode

import (
	"sort"
)

func helper(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			return i - 1
		}
	}
	return -1
}

func swapAtIdx(nums []int, idxA, idxB int) {
	nums[idxA], nums[idxB] = nums[idxB], nums[idxA]
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	swapIdx := helper(nums)
	if swapIdx == -1 {
		sort.Sort(sort.IntSlice(nums))
	} else {
		sort.Sort(sort.IntSlice(nums[swapIdx+1:]))
		for i := swapIdx + 1; i < len(nums); i++ {
			if nums[i] > nums[swapIdx] {
				swapAtIdx(nums, swapIdx, i)
				break
			}
		}
	}
}
