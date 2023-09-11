package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// an string like ("10+4-1") // result in integer

func main() {
	s := "10+4-1"
	fmt.Println(calc(s))
}

func calc(s string) int {
	signs := "+-/%"

	// T = O(N)
	// S = O(1)
	var lastSign string
	var lastDigit string
	var result int
	for _, v := range s {

		result = calculate(&result, &lastDigit, &lastSign)

		if sign := strings.Index(signs, string(v)); sign != -1 {
			lastSign = string(v)

		} else {
			lastDigit += string(v)
		}
	}

	return calculate(&result, &lastDigit, &lastSign)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}

	return i
}

func calculate(result *int, lastDigit, lastSign *string) int {

	if *lastDigit != "" && *lastSign != "" {

		switch *lastSign {

		case "+":
			*result += strToInt(*lastDigit)
			*lastDigit = ""

		case "-":
			*result -= strToInt(*lastDigit)
			*lastDigit = ""

		case "/":
			*result /= strToInt(*lastDigit)
			*lastDigit = ""
		}
	}

	return *result
}
