package main

import "fmt"

func main() {

	// Default type of integer is int
	var a = 10

	// If you want to declare a variable and assign it the zero value
	// You can keep the type and drop the =
	var b int

	// Declaring multiple variables at once with var
	var x, y int = 10, 20

	// You can declare all zero values of the same type
	var k, j int

	// Or of different types
	var g, o = 10, "hello"

	// Declaring multiple variables

	var (
		q    int
		w           = 20
		e    int    = 30
		r, t        = 40, "hello"
		f, h string = "agon", "ademaj"
	)

	fmt.Println(q, w, e, r, t, f, h)

	// When you are within a function you can use := operator
	// to replace var declaration that uses type interface

	// var m = 10
	// m := 10

	// As with var, you can declare multiple variables at once using :=
	// var z, b, c = 10, 3, "hello"
	// z, b, c := 10, 3, "hello"

	fmt.Println(x, y)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(k, j)
	fmt.Println(g, o)

	// The := operator can do one trick that you cannot do with var
	// It allows you to assign values to existing variables too
	// As long as at least one new variable is on the left hand side :=
	// any of the other variables can already exist

	l := 10
	l, p := 30, "hello"

	fmt.Println(l, p)

	// Using := has one limitation
	// If declaring a variable at the package level, you must use var because := is not legal outside of functions

	// In some situations within functions, you should avoid :=

	// You should rarely declare variables outside of functions
	// in what's called the package block

	// As a general rule, you should declare variables in the package block that are effectively immutable

	// Avoid declaring variables outside of functions because they complicate the data flow analysis
}
