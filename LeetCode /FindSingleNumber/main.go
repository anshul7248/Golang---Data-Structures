package main

import "fmt"

func main() {
	var n int
	result := 0
	fmt.Println("Enter the number of elements")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the values of array")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Current array is--->", arr)

	// for _, num := range arr {   // Method 1
	// 	result ^= num
	// }

	for i := 0; i <= len(arr)-1; i++ { // Method 2
		result ^= arr[i]
	}
	fmt.Println("result--->>>", result)
}
