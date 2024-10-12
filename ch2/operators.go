package main

import "fmt"

func main() {
	var a, b int = 10, 3

	fmt.Println("Arithmetic Operators:")
	fmt.Println("a + b =", a+b) // Addition
	fmt.Println("a - b =", a-b) // Subtraction
	fmt.Println("a * b =", a*b) // Multiplication
	fmt.Println("a / b =", a/b) // Division

	// To get a floating-point result
	var c, d int = 10, 3
	result := float64(c) / float64(d)

	fmt.Println("\nConverting to Floating Point for Division")
	fmt.Println("Floating-point division: c / d =", result)

	var x int = 10
	x *= 2
	fmt.Println("\nArithmetic Operators:")
	fmt.Println("x *= 2:", x)

	x += 5
	fmt.Println("x += 5:", x)

	x -= 3
	fmt.Println("x -= 3:", x)

	x /= 2
	fmt.Println("x /= 2:", x)

	x %= 4
	fmt.Println("x %= 4:", x)

	var k, j int = 10, 3
	fmt.Println("\nComparison Operators:")

	fmt.Println("k == j:", k == j) // Equal
	fmt.Println("k != j:", k != j) // Not equal
	fmt.Println("k > j:", k > j)   // Greater than
	fmt.Println("k >= j:", k >= j) // Greater than or equal to
	fmt.Println("k < j:", k < j)   // Less than
	fmt.Println("k <= j:", k <= j) // Less than or equal to

	fmt.Println("\nBitwise Operators")
	var g, o int = 10, 3

	fmt.Println("g & o =", g&o)   // Bitwise AND
	fmt.Println("g | o =", g|o)   // Bitwise OR
	fmt.Println("g ^ o =", g^o)   // Bitwise XOR
	fmt.Println("g &^ o =", g&^o) // Bitwise AND NOT
}
