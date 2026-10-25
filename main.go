package main

import (
	"fmt"
)

func main() {
	fmt.Println(minOperations("abcd"))
	fmt.Println(minOperations("dcba"))
	fmt.Println(minOperations("abdc"))
	fmt.Println(minOperations("dbca"))
}
