package main

func minInsertions(s string) int {
	res := 0
	count_left := 0
	left, right := 0, 0
	for left < len(s) {
		if s[left] == '(' {
			count_left++
			left++
		} else {
			right = left
			for right+1 < len(s) && s[right+1] == ')' {
				right++
			}
			temp := right - left + 1
			if temp%2 != 0 {
				temp++
				res++
			}
			if temp/2 > count_left {
				res += temp/2 - count_left
				count_left = 0
			} else {
				count_left -= temp / 2
			}
			left = right + 1
		}
	}
	if count_left > 0 {
		res += count_left * 2
	}
	return res
}
