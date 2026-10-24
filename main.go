package main

import (
	"fmt"
)

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31"))
}
