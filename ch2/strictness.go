package main

import "fmt"

func main() {
	var x int = 0
	var y int = 5
	var s string = ""

	// Instead of using x or y directly in condition (like in JavaScript)
	// we must explicitly compare the values

	if x == 0 {
		fmt.Println("x is zero")
	}

	if y != 0 {
		fmt.Println("y is non-zero")
	}

	if s == "" {
		fmt.Println("s is an empty string")
	}

	// You cannot do something like "if x" or "if s" directly as in JS
	// Go enforces using comparisons to maintain clarity and avoid "truthy" confusion
}
