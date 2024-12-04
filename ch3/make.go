package main

import "fmt"

// We already see two ways of declaring a slice, using literal or the nil zero value
// While useful, neither way allows us to create an empty slice that already has a length specified.

// That is the job of the build in function called make

func main() {
	x := make([]int, 5)

	// This creates an int slice with a length of 5 and a capacity of 5.
	// So x[0] through x[4] are valid elements, and they are all initialized to 0

	fmt.Println(x)

	// Common beginner mistake is to try to populate those initial elements using append
	a := make([]int, 5)
	a = append(a, 10)

	fmt.Println(a)

	// The 10 is placed at the end of the slice after the zero values in elements 0-4 because
	// append always increases the length of the slice

	// [0, 0, 0, 0, 0, 10]
	// With a length of 6 and a capacity of 10 (capacity was doubled as soon as the sixth element was appended)

	b := make([]int, 6, 10)
	fmt.Println(b)
	// This creates an int slice with a length of 6 and a capacity of 10
}
