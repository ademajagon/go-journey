package main

import "fmt"

func main() {
	fmt.Println("Numeric Types:")

	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	var ua uint8 = 255 // Usually called a byte
	var ub uint16 = 65535
	var uc uint32 = 4294967295
	var ud uint64 = 18446744073709551615

	// Int -> represents signed integers
	// Can hold both positive and negative values
	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)

	// UInt -> represents unsigned integers
	// Can hold only non-negative (positive) values
	fmt.Println("uint8:", ua)
	fmt.Println("uint16:", ub)
	fmt.Println("uint32:", uc)
	fmt.Println("uint64:", ud)

	// Overflow Behavior
	var test int8 = 127
	fmt.Println("Before Overflow:", test)

	test += 1
	fmt.Println("After Overflow:", test) // -128
}
