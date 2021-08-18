package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6}

	n := 3

	for i, v := range array {
		if v == n {
			fmt.Printf("value is: %d and the index is: %d", v, i)
			break
		}
	}
}
