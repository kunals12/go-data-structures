package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))

	language1, language2, language3 := getLanguages()
	fmt.Println(language1, language2, language3)
}

func add(a, b int) int { // a and b are parameters of type int
	return a + b
}

// multiple return values
func getLanguages() (string, string, string) {
	return "golang", "python", "javascript"
}
