package main

import "fmt"

func main() {
	// Explicit type declaration
	var a int = 10

	// Implicit type declaration
	var b string
	b = "Hello, Go!"

	// Short variable declaration
	var c = 20.5

	// Another short variable declaration
	d := "Short declaration"

	// Multiple variable declaration
	var e, f int = 1, 2

	// Print the variables
	fmt.Println("Integer a:", a)
	fmt.Println("String b:", b)
	fmt.Println("Float c:", c)
	fmt.Println("String d:", d)
	fmt.Println("Integers e and f:", e, f)
}
