package main

import (
	"sort"
)

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	k := int64(k1 + k2)
	n := len(nums1)

	sum := int64(0)
	for i := 0; i < n; i++ {
		if nums1[i] > nums2[i] {
			nums1[i] = nums1[i] - nums2[i]
		} else {
			nums1[i] = nums2[i] - nums1[i]
		}
		sum += int64(nums1[i])
	}
	if sum <= k {
		return 0
	}

	sort.Slice(nums1, func(i, j int) bool { return nums1[i] > nums1[j] })
	nums1 = append(nums1, 0)

	for i := 1; i <= n; i++ {
		cost := int64(nums1[i-1]-nums1[i]) * int64(i)
		if cost > k {
			q, r := k/int64(i), k%int64(i)
			hi := int64(nums1[i-1]) - q
			ans := hi*hi*(int64(i)-r) + (hi-1)*(hi-1)*r
			for _, x := range nums1[i:n] {
				ans += int64(x) * int64(x)
			}
			return ans
		}
		k -= cost
	}
	return 0
}
