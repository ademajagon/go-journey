package main

import "fmt"

func main() {
	var x []int
	x = append(x, 10) // assign result to the variable that's passed in

	fmt.Println(x)

	// append function takes at least two parameters
	// a slice of any type and a value of that type

	var a = []int{1, 2, 3}
	a = append(a, 4)
	fmt.Println(a)

	// We can append more than one value
	a = append(a, 5, 6, 7)
	fmt.Println(a)

	// One slice is appended onto another by using the ... operator
	b := []int{10, 20, 30}
	a = append(a, b...)
	fmt.Println(a)

	// It's compile time error if you forget to assign the value from append

	// GO is a call-by-value language.
	// Every time you pass a parameter to a function, Go makes a copy of the value that's passed in.
	// Passing a slice to the append function actually passed a copy of the slice to the function.
	// You then assign the returned slice back to the variable in the calling function.
}
