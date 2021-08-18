package main

import "fmt"

func main() {
	arr := []int{5, 3, 7, 1, 5, 8, 2, 6, 7}

	for i := 0; i < len(arr); i++ {
		mainIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[mainIndex] {
				mainIndex = j
			}
		}

		keepFirstIndex := arr[i]
		arr[i] = arr[mainIndex]
		arr[mainIndex] = keepFirstIndex
	}

	fmt.Println(arr)
}
