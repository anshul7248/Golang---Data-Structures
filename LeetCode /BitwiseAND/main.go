// This program is used find the & of two numbers

// Example like 3 & 5 the output o 3 & 5 will be 1. Let explore the program

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int
	var n2 int

	fmt.Println("Enter first number---->>>.")
	fmt.Scan(&n1)

	fmt.Println("Enter second number---->>>")
	fmt.Scan(&n2)

	binaryn1 := strconv.FormatInt(int64(n1), 2)
	binaryn2 := strconv.FormatInt(int64(n2), 2)

	fmt.Println("Binary of n1 is ---->", binaryn1)
	fmt.Println("Binary of n2 is ---->", binaryn2)

	output := n1 & n2

	fmt.Println("Output of n1-", n1, "&", "n2- ", n2, " is ", output)

}
