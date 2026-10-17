package main

import (
	"strconv"
	"strings"
)

func maxSumOfSquares(num int, sum int) string {
	if num*9 < sum {
		return ""
	}
	res := ""
	res += strings.Repeat("9", sum/9)
	if sum > sum/9*9 {
		res += strconv.Itoa(sum - sum/9*9)
	}
	res += strings.Repeat("0", num-len(res))
	return res
}
