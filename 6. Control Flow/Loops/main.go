package main

import "fmt"

func main() {
	// Simple for loop
	for i := 0; i < 10; i++ {
		fmt.Printf("num  %d ", i)
	}
	fmt.Println()

	// Nested for loop
	// Multiplication table from 1 to 9
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// Loop Control Statements
	// Break statement
	for i := 0; i < 9; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("Break num %d\n", i)
	}
	// Continue statement
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Continue num %d\n", i)
	}
	// GoTo statement
	for i := 0; i < 9; i++ {
		fmt.Printf("GoTo num %d\n", i)
		if i == 3 {
			goto end
		}
	}
end:
	fmt.Println("Loop Ended.")
	// Fun loop
	test := 0
loop:
	if test < 8 {
		fmt.Println("Num", test)
		test++
		goto loop
	}

	//	Infinite loop
	//	for {
	//		fmt.Println("Infinite Loop")
	//	}
	//	Using goto for infinite loop
	//
	// start:
	//
	//	fmt.Println("Infinite Loop.")
	//	goto start
}
