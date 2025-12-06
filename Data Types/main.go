package main

import (
	"fmt"
	"unsafe"
)

func integerTypes() {
	// Declare a variable of type signed int
	var si int
	var si8 int8
	var si16 int16
	var si32 int32
	var si64 int64

	// Declare a variable of type unsigned int
	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64

	// Print the sizes of the variables
	fmt.Printf("Size of signed int: %d bytes\n", unsafe.Sizeof(si))
	fmt.Printf("Size of signed int8: %d bytes\n", unsafe.Sizeof(si8))
	fmt.Printf("Size of signed int16: %d bytes\n", unsafe.Sizeof(si16))
	fmt.Printf("Size of signed int32: %d bytes\n", unsafe.Sizeof(si32))
	fmt.Printf("Size of signed int64: %d bytes\n\n", unsafe.Sizeof(si64))

	// Print the sizes of the unsigned variables
	fmt.Printf("Size of unsigned int: %d bytes\n", unsafe.Sizeof(ui))
	fmt.Printf("Size of unsigned int8: %d bytes\n", unsafe.Sizeof(ui8))
	fmt.Printf("Size of unsigned int16: %d bytes\n", unsafe.Sizeof(ui16))
	fmt.Printf("Size of unsigned int32: %d bytes\n", unsafe.Sizeof(ui32))
	fmt.Printf("Size of unsigned int64: %d bytes\n", unsafe.Sizeof(ui64))
}

func floatTypes() {
	var f32 float32
	var f64 float64
	fmt.Printf("Size of low precision float32: %d bytes\n", unsafe.Sizeof(f32))
	fmt.Printf("Size of high precision float64: %d bytes\n\n", unsafe.Sizeof(f64))

	var lowPrecision float32 = 1.0123456789012345
	var highPrecision float64 = 1.0123456789012345
	fmt.Printf("Low Precision float32 value: %f\n", lowPrecision)
	fmt.Printf("High Precision float64 value: %f\n", highPrecision)
}

func booleanTypes() {
	var active bool = true
	var notActive bool = false

	fmt.Printf("Is Active: %t\n", active)
	fmt.Printf("Is not Active: %t\n", notActive)
}

func complexTypes() {
	var cn64 complex64 = complex(5, 10)
	var cn128 complex128 = complex(20, 30)

	fmt.Printf("Complex Number (complex64): %v\n", cn64)
	fmt.Printf("Complex Number (complex128): %v\n", cn128)
}

func stringTypes() {
	var str string = "Mamadreza"
	fmt.Printf("My name is %s\n", str)
}

func main() {
	// Call the function to display sizes of integer types
	fmt.Println("Sizes of Integer Types:")
	integerTypes()
	fmt.Printf("End of Integer Types\n")

	// Call the function to display sizes of float types
	fmt.Println("\nSizes of Float Types:")
	floatTypes()
	fmt.Printf("End of Float Types\n")

	// Call the function to demonstrate boolean types
	fmt.Println("\nBoolean Types:")
	booleanTypes()
	fmt.Printf("End of Boolean Types\n")

	// Call the function to demonstrate complex number types
	fmt.Println("\nComplex Number Types:")
	complexTypes()
	fmt.Printf("End of Complex Number Types\n")

	// Call the function to demonstrate string types
	fmt.Println("\nString Types:")
	stringTypes()
	fmt.Printf("End of String Types\n")
}
