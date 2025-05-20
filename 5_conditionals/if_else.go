package main

import "fmt"

func main() {
	i := 1

	if i == 2 {
		fmt.Println("i is 2")
	} else if i == 1 {
		fmt.Println("i is 1")
	} else {
		fmt.Println("i is not found")
	}
}
