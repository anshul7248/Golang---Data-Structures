package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)

	if n&1 == 0 {
		fmt.Println(n, " is even")

	} else {
		fmt.Println(n, "is odd")
	}
}
