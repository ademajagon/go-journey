package main

import (
	"fmt"
	"math"
)

// 1. Write a program that declares an integer variable called i with the value 20
// Assign i to a floating-point variable named f. Print out i and f

/*func main() {
	var i int = 20

	var f float64 = float64(i)

	fmt.Println(i, f)
}
*/

// 2. Write a program that declares a constant called value that can be assigned to both an integer and a floating-point variable
// Assign it to an integer called i and a floating-point variable called f. Print out i and f

/*func main() {
	const value = 20

	var i int = value
	var f float64 = value

	fmt.Println(i, f)
}*/

// 3. Write a program with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64
// Assign each variable the maximum legal value for its type; then add 1 to each variable. Print out their values

func main() {
	var b uint8 = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
