package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)

	result := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count := 0
		num := i
		// lastbit:= num & 1
		for num > 0 {
			lastbit := num & 1
			count = count + lastbit
			num = num >> 1
		}

		result[i] = count
	}
	fmt.Println("Bit count for each number from 0 to", n, ":", result)

}
