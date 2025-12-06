package main

import "fmt"

func main() {
	const AGE int = 30
	const GREETING string = "Hello, World!"
	const PI float64 = 3.14

	fmt.Printf("%s I am %d years old. Pi is approximately %.2f.\n", GREETING, AGE, PI)

	const (
		DECIMAL     = 255
		HEXADECIMAL = 0xFF
		OCTAL       = 0377
		BINARY      = 0b11111111
	)
	fmt.Printf("Decimal: %d, Hexadecimal: %d, Octal: %d, Binary: %d\n", DECIMAL, HEXADECIMAL, OCTAL, BINARY)

	const AVAGADRO float64 = 6.022e23
	fmt.Printf("Avogadro's number is approximately %.3e\n", AVAGADRO)

	 
}
