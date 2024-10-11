package main

import "fmt"

func main() {
	decimal := 42      // base 10
	binary := 0b101010 // base 2
	octal := 052       // base 8
	hex := 0x2A        // base 16
	fmt.Println("Integer Literals:")
	fmt.Println("Decimal:", decimal)
	fmt.Println("Binary:", binary)
	fmt.Println("Octal:", octal)
	fmt.Println("Hex:", hex)

	decimalFloat := 3.14
	scientificFloat := 1.5e2
	hexFloat := 0x1.2p3
	fmt.Println("\nFloating-Point Literals:")
	fmt.Println("Float:", decimalFloat)
	fmt.Println("Scientific:", scientificFloat)
	fmt.Println("Hex:", hexFloat)

	r1 := 'A'
	r2 := '\n'
	r3 := '\u03A9'
	fmt.Println("\nRune Literals:")
	fmt.Printf("Character: %c\n\n", r1)
	fmt.Printf("Character: %c\n\n", r2) // This will print a new line
	fmt.Printf("Character: %c\n\n", r3)
}
