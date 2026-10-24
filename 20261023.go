package main

func sortPermutation(nums []int) int {
	l := len(nums)
	res := 0
	tag := true
	for i := 0; i < l; i++ {
		if nums[i] == i {
			continue
		}
		if tag {
			tag = false
			res = i & nums[i]
		} else {
			res &= i & nums[i]
		}
	}
	return res
}
