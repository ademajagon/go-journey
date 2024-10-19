package main

import "fmt"

func main() {
	// Creates an array of three integers
	// If no values specified, all elements are initialized to the zero value
	var x [3]int
	fmt.Println(x)

	// If you have initial values for the array
	// Specify them with an array literal
	var y = [3]int{2, 4, 4}
	fmt.Println(y)

	// If having a sparse array (an array where most of the values are set to zero)
	// Specify the indices with nonzero values in the array literal
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z)

	// When using an array literal to init an array
	// You can replace the number that specifies the number of elements in the array with ...
	var a = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)

	// You can use != or == to compare two arrays.
	// Arrays are equal if they are the same length and contain equal values.
	var b = [...]int{3, 2, 1}
	var c = [3]int{3, 2, 1}
	fmt.Println(b == c) // prints true
	fmt.Println(b != c) // print false

	// Go has only one-dimensional arrays, but you can stimulate multidimensional arrays
	var d [3][2]int
	fmt.Println(d)

	// Like most languages, arrays in GO are read and written using brackets syntax
	var g = [3]int{3, 4, 1}
	g[0] = 10
	fmt.Println(g[0])
	fmt.Println(g[2])

	// You cannot read or write past the end of an array or use a negative index
	// If you do this with a constant or literal index, it's a compile time error.

	// An out-of-bounds read or write with a variable index compiles but fails at runtime with a panic.

	// The built-in function len takes in an array and returns its length.
	fmt.Println(len(g))

	// Arrays in GO are rarely used explicitly. This is because they come with an unusual limitation
	// GO considers the size of the array to be part of type of the array.

	// This makes an array that's declared to be [3]int a different type from an array that's declared to be [4]int.
	// This also means that you cannot use a variable to specify the size of an array,
	// because type must be resolved at compile-time, not at runtime.

	// var arr1 [3]int
	// var arr2 [4]int
	// arr1 = arr2 // This will result in compile-time error.

	// You can't use a type conversion to directly convert arrays of different sizes to identical types.
	// Because, you can't convert arrays of different sizes into each other, you can't write a function that works with arrays of any size
	// And, you can't assign arrays of different sizes to the same variable.

	// Because of these restrictions, don't use arrays unless you know the exact length you need ahead of time.

	// This raises the question, why is such a limited feature in the language?
	// The main reason arrays exist in GO, is to provide the backing store for slices, which are one of the most useful features of Go.
}
