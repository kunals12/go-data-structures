/*
Go doesn’t have built-in enum types like languages such as Java or C#, but you can simulate enums using constants and custom types.
Below is a complete, beginner-friendly explanation with:

  - Enum simulation using iota
  - Custom string representations
  - Real-world use case
  - Code with comments

✅ What Are Enums in Go?
In Go, enums are typically represented using:

  - A custom type (type Status string or type Status int)
  - A set of constants using const and optionally iota for integer enums
*/
package main

import (
	"fmt"
)

// Define a new type 'OrderStatus' as a string
type OrderStatus string

// Define possible enum values as constants
const (
	StatusPending   OrderStatus = "Pending"
	StatusConfirmed OrderStatus = "Confirmed"
	StatusShipped   OrderStatus = "Shipped"
	StatusDelivered OrderStatus = "Delivered"
	StatusCancelled OrderStatus = "Cancelled"
)

// Order struct with status field using the custom enum type
type Order struct {
	id     uint64
	amount uint32
	status OrderStatus // enum type
}

// Function to print readable order
func (o Order) printStatus() {
	fmt.Printf("Order #%d has status: %s\n", o.id, o.status)
}

func main() {
	order := Order{
		id:     1001,
		amount: 499,
		status: StatusPending,
	}

	order.printStatus()

	// Update status using enum
	order.status = StatusConfirmed
	order.printStatus()

	order.status = StatusDelivered
	order.printStatus()
}

/*
type Role int

const (
	Admin Role = iota
	Editor
	Viewer
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Editor:
		return "Editor"
	case Viewer:
		return "Viewer"
	default:
		return "Unknown"
	}
}

// Tip: Validating Enum Values
func IsValidStatus(status OrderStatus) bool {
	switch status {
	case StatusPending, StatusConfirmed, StatusShipped, StatusDelivered, StatusCancelled:
		return true
	default:
		return false
	}
}

*/
