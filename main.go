package main

import (
	"fmt"
)

func main() {
	fmt.Println(sortPermutation([]int{0, 3, 2, 1}))
	fmt.Println(sortPermutation([]int{0, 1, 3, 2}))
	fmt.Println(sortPermutation([]int{3, 2, 1, 0}))
	fmt.Println(sortPermutation([]int{0, 3, 1, 2}))
}
