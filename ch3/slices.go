package main

import (
	"fmt"
	"slices"
)

func main() {
	// When you want a data structure that holds a sequence of values, a slice is what you should use
	// What makes slices so useful is that you can grow slices as needed.
	// This is because the length of a slice is NOT part of its type.
	// This removes the biggest limitations of arrays and allows you to write a single function that processes slices of any size

	// Subtle differences between arrays and slices exists.
	// You don't specify the size of the slice when you declare it.

	var x = []int{10, 20, 30}
	fmt.Println(x)

	// Using [...] makes an array. Using [] makes a slice.

	// Just as with arrays you can also specify only the indices with nonzero values in the slice literal
	var y = []int{1, 5: 4, 3, 4, 10: 3, 3}
	fmt.Println(y)

	// Can also stimulate multidimensional slices and make a slice of slices
	var z [][]int
	fmt.Println(z)

	//var a = []int{1, 2, 3}
	//var b = []int{1, 2, 3}

	// It is compile-time error to use == to see if two slices are identical or != to see if different.
	//fmt.Println(a == b)

	// The only thing you can compare a slice with using == is nil
	// A nil slice contains nothing.

	var c []int
	fmt.Println(c == nil)

	// In Go 1.21, slices package in the standard library includes two functions to compare slices

	// slices.Equal function takes in two slices and returns true if the slices are the same length and all elements are equal.
	// slices.EqualFunc, lets you pass in a function to determine equality and does not require the slice elements to be comparable.

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	d := []int{6, 7, 8, 9, 10}
	//e := []string{"a", "g", "o", "n"}

	fmt.Println(slices.Equal(a, b)) // print true
	fmt.Println(slices.Equal(b, d)) // print false
	//fmt.Println(slices.Equal(d, e)) // does not compile

	// reflect package contains a function called DeepEqual that can compare almost anything, including slices.
	// It's a legacy function, primarily intended for testing.
	// Before the inclusion of slices.Equal and slices.EqualFunc, reflect.DeepEqual was often used to compare slices.
	// Don't use it as it is slower and less safe than using the functions in the slices package.

	// Len
	// Passing a nil slice to len returns 0
}
