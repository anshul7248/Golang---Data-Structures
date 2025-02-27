package main

import "fmt"

func main() {
	var n int32
	var result int32 = 0
	fmt.Println("Enter the number")
	fmt.Scan(&n)

	for n != 0 {
		lastdigit := n % 10
		result = result*10 + lastdigit
		n = n / 10
	}

	fmt.Println("Reversed output", result)
}
