package main

import "fmt"

func main() {
	age := 14
	if age < 18 {
		fmt.Println("Not allowed to drink alcohol.")
	} else {
		fmt.Println("Allowed to drink alcohol.")
	}

	score := 93
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		if score >= 85 {
			fmt.Println("Grade: B+")
		} else {
			fmt.Println("Grade: B")
		}
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
}
