package main

func sumOfSquares(nums []int) int {
	res := 0
	l := len(nums)
	for i := 0; i < l; i++ {
		if l%(i+1) == 0 {
			res += nums[i] * nums[i]
		}
	}
	return res
}
