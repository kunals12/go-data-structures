/*
An interface defines a set of method signatures.
Any type that implements those methods implicitly satisfies the interface — no explicit declaration needed.

✅ Real-world analogy: An interface is like a contract.
If someone can perform the required actions (methods), they fulfill the role, regardless of their actual type.
*/

package main

import (
	"fmt"
	"time"
)

// Interface definition
type OrderProcessor interface {
	changeStatus(status string)
	printDetails()
}

// Order struct
type Order struct {
	id        uint64
	amount    uint32
	status    string
	createdAt time.Time
}

// Constructor for Order
func newOrder(id uint64, amount uint32, status string) *Order {
	return &Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

// Pointer receiver: modifies original struct
func (o *Order) changeStatus(status string) {
	o.status = status
}

// Value receiver: prints details (doesn't need to modify anything)
func (o Order) printDetails() {
	fmt.Printf("Order ID: %d | Amount: %d | Status: %s | Created At: %s\n",
		o.id, o.amount, o.status, o.createdAt.Format(time.RFC1123))
}

// Another struct implementing the same interface
type Subscription struct {
	userID    string
	plan      string
	status    string
	startedAt time.Time
}

func newSubscription(userID, plan string) *Subscription {
	return &Subscription{
		userID:    userID,
		plan:      plan,
		status:    "Active",
		startedAt: time.Now(),
	}
}

func (s *Subscription) changeStatus(status string) {
	s.status = status
}

func (s Subscription) printDetails() {
	fmt.Printf("Subscription for User: %s | Plan: %s | Status: %s | Started At: %s\n",
		s.userID, s.plan, s.status, s.startedAt.Format(time.RFC1123))
}

func process(p OrderProcessor, status string) {
	p.changeStatus(status) // Interface method call
	p.printDetails()       // Interface method call
}

func main() {
	order := newOrder(101, 499, "Pending")
	sub := newSubscription("user_123", "Pro")

	fmt.Println("Processing Order:")
	process(order, "Confirmed") // Order implements OrderProcessor

	fmt.Println("\nProcessing Subscription:")
	process(sub, "Cancelled") // Subscription implements OrderProcessor
}

/*
🔸 Interface:

type OrderProcessor interface {
	changeStatus(status string)
	printDetails()
}
This defines behavior — any type with both changeStatus() and printDetails() satisfies the interface.

🔸 Implicit Implementation:
Neither Order nor Subscription declares that it implements OrderProcessor. Go automatically sees that both types have the required methods.

🔸 Method Sets:
A pointer receiver (*Order) satisfies an interface requiring that method.

If the method has a pointer receiver, you must pass a pointer to the interface, not a value.
*/
