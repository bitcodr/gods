package main

import "fmt"

func main() {
	arr := []int{5, 3, 1, 8, 3, 7} // 3 1 5 3 7 8 - 1 3 3 5 7 8
	fmt.Println(sort(arr))
}

// T -> O(N ^ 2)
// S -> O(1)
func sort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}

	return arr
}
