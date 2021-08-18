package main

import "fmt"

func main() {
	n := 12345

	var x int

	for n > 0 {
		remainder := n % 10
		x *= 10
		x += remainder
		n /= 10
	}

	fmt.Println(x)
}
