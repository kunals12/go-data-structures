/*
🔹 Pointer:
A pointer is a variable that stores the memory address of another variable.

In func changeNum(num *int), *int means a pointer to an int.

🔹 Reference (& operator):
The & operator is used to get the address of a variable.

changeNum(&num) passes the reference (address) of num.

🔹 Dereference (* operator):
The * operator is used to access the value stored at a memory address.

*num = 5 means we are updating the value at the referenced memory location.
*/

package main

import "fmt"

// changeNum takes a pointer to an int as a parameter
// This means it receives the *address* of the variable, not a copy of its value
func changeNum(num *int) {
	// *num is the dereferenced value — it allows us to modify the original variable
	*num = 5 // Dereferencing the pointer to change the actual value at that memory address
}

func main() {
	num := 1 // Declare an integer variable 'num' with a value of 1

	// Print the original value of num
	fmt.Println("Before:", num)

	// Pass the address of num using &num (reference)
	// This allows changeNum to modify the original variable
	changeNum(&num)

	// Print the modified value of num
	fmt.Println("After:", num)
}

/*
💡 Real-World Use Case: Modify Shared Data
Suppose you're writing a function to update the price of a product.
If you pass the price by value, the function will modify a copy and the original remains unchanged.
But using a pointer lets you modify the original variable directly.


func updatePrice(price *float64) {
	*price = *price * 0.9 // Apply 10% discount
}

func main() {
	price := 100.0
	updatePrice(&price)
	fmt.Println("Discounted Price:", price) // Output: 90
}

*/
