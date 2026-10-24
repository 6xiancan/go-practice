package main

import (
	"strconv"
)

func _day2Num(dayString string) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	month, _ := strconv.Atoi(dayString[0:2])
	day, _ := strconv.Atoi(dayString[3:5])
	res := 0
	for i := 0; i < month-1; i++ {
		res += months[i]
	}
	res += day
	return res
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	startAlice, endAlice := _day2Num(arriveAlice), _day2Num(leaveAlice)
	startBob, endBob := _day2Num(arriveBob), _day2Num(leaveBob)
	start, end := 0, 0
	if startAlice > startBob {
		start = startAlice
	} else {
		start = startBob
	}
	if endAlice < endBob {
		end = endAlice
	} else {
		end = endBob
	}
	if start <= end {
		return end - start + 1
	}
	return 0
}
