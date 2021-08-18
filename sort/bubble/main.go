package main

import "fmt"

func main() {
	arr := []int{5, 2, 6, 4, 7, 8, 1, 7, 3}

	//2 < 5

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i] = arr[j]
				arr[j] = arr[i]
			}
		}
	}

	fmt.Println(arr)
}
