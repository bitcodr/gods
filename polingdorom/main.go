package main

import (
	"fmt"
)

func main() {
	s := "quyjjdcgsvvsgcdjjyq"

	var left int
	right := len(s) - 1

	result := -1

	for left < right {
		if string(s[left]) != string(s[right]) {
			if string(s[left]) != string(s[right-1]) && string(s[left+1]) == string(s[right]) {
				result = left
				break
			} else {
				result = right
				break
			}
		}

		left++
		right--
	}

	fmt.Println(result)
}
