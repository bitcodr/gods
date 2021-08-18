package main

import (
	"fmt"
	"os"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	n := 4

	response := find(arr, 0, len(arr)-1, n)
	if response != -1 {
		fmt.Println(response)
		os.Exit(0)
	}

	fmt.Println("could not find the key")
}

func find(arr []int, left, right, k int) int {
	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == k {
			return mid
		}

		if k > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
