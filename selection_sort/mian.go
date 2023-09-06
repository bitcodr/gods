package main

import "fmt"

func main() {
	arr := []int{3, 5, 6, 2, 8, 1}
	fmt.Println(sort(arr))
}

// T -> O(N^2)
// S -> O(1) -- because it is in-place
func sort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		smallest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}

		temp := arr[i]
		arr[i] = arr[smallest]
		arr[smallest] = temp
	}

	return arr
}
