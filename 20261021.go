package main

import (
	"sort"
)

func _sortStrings(s string) string {
	rs := []rune(s)
	sort.Slice(rs, func(i, j int) bool { return rs[i] < rs[j] })
	sort_s := string(rs)
	return sort_s
}

func minOperations(s string) int {
	l := len(s)
	if l == 1 {
		return 0
	}
	if l == 2 && s[0] > s[1] {
		return -1
	}
	if l == 2 {
		return 0
	}
	valid_s := _sortStrings(s)
	if s == valid_s {
		return 0
	}
	res := 0
	temp_left, temp_right := _sortStrings(s[0:l-1])+string(s[l-1]), string(s[0])+_sortStrings(s[1:l])
	res++
	if temp_left == valid_s || temp_right == valid_s {
		return res
	}
	temp_left, temp_right = string(temp_left[0])+_sortStrings(temp_left[1:l]), _sortStrings(temp_right[0:l-1])+string(temp_right[l-1])
	res++
	if temp_left == valid_s || temp_right == valid_s {
		return res
	}
	return res + 1
}
