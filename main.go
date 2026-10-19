package main

import (
	"fmt"
)

func main() {
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 2))
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 3))
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 5))
	fmt.Println(targetIndices([]int{1, 2, 5, 2, 3}, 4))
}
