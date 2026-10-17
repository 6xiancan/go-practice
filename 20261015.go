package main

func sumCounts(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		set := make(map[int]bool)
		for j := i; j < len(nums); j++ {
			set[nums[j]] = true
			res += len(set) * len(set)
		}
	}
	return res
}
