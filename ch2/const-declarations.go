package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	//x = x + 1 // this will not compile!
	//y = "bye" // this will nto compile!
	//fmt.Println(x)
	//fmt.Println(y)

	//a := 5
	//b := 10
	//const c = a + b

	// Constants in GO are a way to give names to literals
	// There is no way in GO to declare that a variable is immutable

	// Constants can be typed or untyped
	// A typed constant can be directly assigned only to a variable of that type

	// If giving a name to a mathematical constant that could be used with multiple numeric types, keep the constant untyped

	// Here is what an untyped constant declaration look like
	const a = 10

	// All the following assignments are legal
	var j int = a
	var k float64 = a
	var l byte = a

	fmt.Println(j, k, l)

	// Here is what a typed constant declaration look like
	const typedX int = 10
	// This constant can be assigned directly only to an int.
	// Assigning it to any other type produces a compile-time error like this:
}
