package main

import "sort"

func minSwaps(nums []int) int {
	res := 0
	sort_nums := append([]int(nil), nums...)
	sort.Slice(sort_nums, func(i, j int) bool {
		sum_i, sum_j := 0, 0
		x, y := sort_nums[i], sort_nums[j]
		for x > 0 {
			sum_i += x % 10
			x /= 10
		}
		for y > 0 {
			sum_j += y % 10
			y /= 10
		}
		if sum_i < sum_j {
			return true
		} else if sum_i > sum_j {
			return false
		} else {
			return sort_nums[i] < sort_nums[j]
		}
	})
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] != sort_nums[i] {
			m[nums[i]] = sort_nums[i]
		}
	}
	for len(m) > 0 {
		count := 1
		key := 0
		for temp := range m {
			key = temp
			break
		}
		value := m[key]
		delete(m, key)
		for value != key {
			count++
			temp := m[value]
			delete(m, value)
			value = temp
		}
		res += count - 1
	}
	return res
}
