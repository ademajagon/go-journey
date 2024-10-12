package main

import "fmt"

func main() {

	// GO requirement is that every declared local variable must be read
	// It is a compile-time error to declare a local variable and to not read its value

	x := 10 // this assignment isn't read!
	x = 20

	fmt.Println(x)
	x = 30 // this assignment isn't read!

	// GO compiler won't stop you from creating unread package-level variables
	// This is one more reason you should avoid creating package-level variables.

	// GO compiler allows you to create unread constants with const
	// Constants in GO are calculated at compile time and cannot have any side effects
}
