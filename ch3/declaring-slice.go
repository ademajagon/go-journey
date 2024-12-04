package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data []int
	fmt.Println(data, len(data))

	var x = []int{} // this creates a slice with zero length and zero capacity
	// it's confusingly different from a nil slice
	// Because of implementations reasons, comparing a zero-length slice to nil returns false
	fmt.Println(x, len(x))

	fmt.Println(reflect.DeepEqual(data, x)) // false

	// While comparing a nil slice to nil returns true

	var a []int
	var b []int
	fmt.Println(reflect.DeepEqual(a, b)) // true

	// For simplicity favor nil slices, a zero length slice is useful only when converting a slice to JSON
}
