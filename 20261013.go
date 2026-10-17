package main

import (
	"sort"
)

func maxAlternatingSum(nums []int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < 0 && nums[j] < 0 {
			return -nums[i] > -nums[j]
		} else if nums[i] < 0 && nums[j] >= 0 {
			return -nums[i] > nums[j]
		} else if nums[i] >= 0 && nums[j] < 0 {
			return nums[i] > -nums[j]
		} else {
			return nums[i] > nums[j]
		}
	})
	left, right, res, tag := 0, len(nums)-1, int64(0), true
	for left <= right {
		if tag {
			res += int64(nums[left]) * int64(nums[left])
			left++
		} else {
			res -= int64(nums[right]) * int64(nums[right])
			right--
		}
		tag = !tag
	}
	return res
}
