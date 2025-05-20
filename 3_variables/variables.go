package main

import "fmt"

func main() {
	// String
	var name string = "Golang"
	fmt.Println("Hello", name)

	// Integer
	var age int = 30
	fmt.Println("My age is", age)

	// Boolean
	var isTrue bool = true
	fmt.Println("Is it true?", isTrue)

	// Float
	var pi float64 = 3.14159
	fmt.Println("Pi is", pi)

	// Complex
	var complexNumber complex128 = 3 + 4i
	fmt.Println("Complex number is", complexNumber)

	// Shorthand syntax
	myName := "Kunal"
	fmt.Println("My name is", myName)

	// Constants
	const companyName = "Xyz"
	fmt.Println("Company name is", companyName)

	// let's try to change the value
	// companyName = "Abc" // We can't do this
	// fmt.Println("Company name is", companyName)

	// Multiple variable declaration
	var (
		firstName = "John"
		lastName  = "Doe"
	)
	fmt.Println("Full name is", firstName, lastName)
}
