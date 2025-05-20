package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
}

func add(a, b int) int { // a and b are parameters of type int
	return a + b
}

// multiple return values
func swap(x, y string) (string, string) {
	return y, x
}
