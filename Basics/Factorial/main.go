package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number")
	fmt.Scan(&n)

	result := 1

	for i := n; i >= 1; i-- {
		result = result * i
	}

	fmt.Println("result ", result)
}
