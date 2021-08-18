package main

import "fmt"

//O(n+m)
func main() {
	x := []int{4, 7, 3, 6, 9, 3, 8, 2, 1, 4, 6, 0, 6}
	y := []int{4, 9, 3, 1, 9, 7, 2, 8}

	m := make(map[int]int)

	for i := 0; i < len(x); i++ {
		m[x[i]] = x[i]
	}

	var results []int

	for i := 0; i < len(y); i++ {
		if _, ok := m[y[i]]; ok {
			results = append(results, y[i])
			delete(m, y[i])
		}
	}

	fmt.Println(results)

}
