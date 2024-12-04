package main

import "fmt"

// clear function takes in a slice and sets all of the slice's elements to their zero value.
// The length of the slice remains unchanged.

func main() {
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))

	clear(s)
	fmt.Println(s, len(s))

	// Zero value for a string is "" !
}
