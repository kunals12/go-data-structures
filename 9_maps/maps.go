package main

import "fmt"

func main() {
	m := make(map[int]string)

	m[1] = "Hello"
	m[2] = "World"

	fmt.Println(m)
	fmt.Println(m[1])
	fmt.Println(len(m)) // prints length of map

	// delete
	delete(m, 1) // deletes key 1
	fmt.Println(m)

	// check if key exists
	_, ok := m[1]
	fmt.Println(ok)

	// clear map
	clear(m)
	fmt.Println(m)

	// another way to initialize map
	m2 := map[int]string{1: "Hello", 2: "World"}
	fmt.Println(m2)
}
