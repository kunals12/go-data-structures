package main

import (
	"fmt"
	"time"
)

// Order struct represents a simple order entity
type Order struct {
	id        uint64    // Unique identifier for the order
	amount    uint32    // Total amount of the order
	status    string    // Current status of the order (e.g., Received, Confirmed)
	createdAt time.Time // Timestamp when the order was created (in nanoseconds)
}

// Constructor function that returns a pointer to a new Order
func newOrder(id uint64, amount uint32, status string) *Order {
	return &Order{ // Returning the address of the newly created struct (reference)
		id:     id,
		amount: amount,
		status: status,
	}
}

// changeStatus is a method with a pointer receiver
// Using *Order allows modification of the original struct instance
func (o *Order) changeStatus(_status string) {
	o.status = _status // Dereferencing happens automatically; status is updated in the original Order
}

// getAmount is a method with a value receiver
// It does NOT modify the original Order, just returns a copy of the amount
func (o Order) getAmount() uint32 {
	return o.amount
}

func main() {
	// Create a new Order struct instance using a struct literal (value type)
	order1 := Order{
		id:     1,
		amount: 100,
		status: "Received",
	}

	// Set creation time using time.Now()
	order1.createdAt = time.Now()

	// Print the order before status change
	fmt.Println("Order before:", order1)

	// Call method with pointer receiver to change status
	// Even though order1 is a value type, Go automatically takes its address when calling pointer methods
	order1.changeStatus("Confirmed")

	// Print the updated order
	fmt.Println("Order after status changed:", order1)

	// Call value receiver method to get the amount
	fmt.Println("Order1 amount:", order1.getAmount())

	// Create another Order using the constructor function that returns a pointer
	order2 := newOrder(2, 299, "Confirmed")
	fmt.Println("Order2:", order2)
}

/*
🧠 Real-World Use Case:
Imagine you are building an e-commerce system. When an order is received, its status is "Pending". As it gets processed, you update the status to "Confirmed", "Shipped", and eventually "Delivered".

Using pointer receivers allows you to modify the actual Order object, instead of working on a copy.


*/
