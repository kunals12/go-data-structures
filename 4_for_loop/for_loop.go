package main

import "fmt"

func main() {
	// For is the only loop which we can use in golang

	// While Loop
	fmt.Println("While Loop")
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	/*
		Infinite Loop

		for {
			fmt.Println("Infinite Loop")
		}
	*/

	// For Loop
	fmt.Println("For Loop")
	for i := 1; i <= 10; i++ {
		// break // We can break loop
		if i == 2 {
			continue // We can continue to skip current iteration
		}
		fmt.Println(i)
	}

	// Range Loop
	fmt.Println("Range Loop")
	for i := range 10 {
		fmt.Println(i)
	}

	for i, v := range "Hello World" { // v is rune
		fmt.Println(i, v)
	}
}
