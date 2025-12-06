package main

import "fmt"

func main() {
	// Arithmetic Operators
	A := 10
	B := 20

	fmt.Println("A:", A)
	fmt.Println("B:", B)

	fmt.Println("A + B:", A+B)
	fmt.Println("A - B:", A-B)
	fmt.Println("A * B:", A*B)
	fmt.Println("A / B:", A/B)
	fmt.Println("A % B:", A%B)
	A++
	fmt.Println("A++:", A)
	B--
	fmt.Println("B--:", B)

	// Relational Operators
	fmt.Println("A == B:", A == B)
	fmt.Println("A != B:", A != B)
	fmt.Println("A > B:", A > B)
	fmt.Println("A < B:", A < B)
	fmt.Println("A >= B:", A >= B)
	fmt.Println("A <= B:", A <= B)

	// Logical Operators
	X := true
	Y := false

	fmt.Println("X && Y:", X && Y)
	fmt.Println("X || Y:", X || Y)
	fmt.Println("!X:", !X)
	fmt.Println("!Y:", !Y)

	// Assignment Operators
	A = 10
	B = 20

	A += B
	fmt.Println("A += B", A)
	A -= B
	fmt.Println("A -= B", A)
	A *= B
	fmt.Println("A *= B", A)
	A /= B
	fmt.Println("A /= B", A)
	A %= B
	fmt.Println("A %= B", A)

	C := 30

	// Bitwise Operators
	// Bitwise AND
	fmt.Println("C & 2:", C&2)
	// Bitwise OR
	fmt.Println("C | 2:", C|2)
	// Bitwise XOR
	fmt.Println("C ^ 2:", C^2)
	// Bitwise Left Shift
	fmt.Println("C << 1:", C<<1)
	// Bitwise Right Shift
	fmt.Println("C >> 1:", C>>1)
	// Bitwise NOT
	fmt.Println("~C:", ^C)

	// Miscellaneous Operators
	// Address Operator
	ptr := &A
	fmt.Println("Address of D:", ptr)
	// Pointer Operator
	fmt.Println("Value at pointer ptr:", *ptr)
}
