package main

import "sort"

func targetIndices(nums []int, target int) []int {
	sort.Ints(nums)
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		if nums[i] == target {
			res = append(res, i)
		}
	}
	return res
}
