package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	var result int

	fmt.Println("Enter the number")
	fmt.Scan(&num)

	binary := strconv.FormatInt(int64(num), 2)
	fmt.Println("Binary of entered number is ----->>>>", binary)
	for num != 0 {
		lastbit := num & 1
		if lastbit == 1 {
			result++
		}
		num = num >> 1
	}
	fmt.Println("Number of 1 bits are--->>>", result)
}
