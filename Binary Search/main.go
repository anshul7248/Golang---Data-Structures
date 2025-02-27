package main

import "fmt"

func main() {

	arr := []int{20, 52, 63, 69, 70, 85}

	start := 0
	end := len(arr) - 1
	target := 70

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			fmt.Println("Element is found on index ----->>>", mid)
			break
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
}
