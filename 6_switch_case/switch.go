package main

import (
	"fmt"
	"time"
)

func main() {
	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		case bool:
			fmt.Println("I'm a bool")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whoAmI(21)
	whoAmI("Kunal")
	whoAmI(true)
}
