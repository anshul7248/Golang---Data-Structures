package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Println("Enter the number to check the power of 2")
	fmt.Scan(&num)

	binary := strconv.FormatInt(int64(num), 2)
	fmt.Println("Binary    ", binary)
	count := 0
	for _, v := range binary {
		if v == '1' {
			count++
		}
	}
	if count == 1 {
		fmt.Println("It is a power of 2")
	} else {
		fmt.Println("It is not the power of 2")
	}
}
