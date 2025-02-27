package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)

	binary := strconv.FormatInt(int64(n), 2)

	fmt.Println("Binary of entered number is--->", binary)

	bitlength := len(binary)

	mask := (1 << bitlength) - 1

	complement := ^n & mask

	complementBitMask := strconv.FormatInt(int64(complement), 2)

	fmt.Println("Bitwise complement in decimal:", complement)
	fmt.Println("Bitwise complement in binary:", complementBitMask)
}
