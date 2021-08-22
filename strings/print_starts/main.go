package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 4

	char := "*"

	x := 1

	//O(nm)

	for n > 0 {
		fmt.Printf("%"+strconv.Itoa(n)+"v", " ")
		for i := 0; i < x; i++ {
			fmt.Printf("%s ", char)
		}
		fmt.Println("")
		x++
		n--
	}

}
