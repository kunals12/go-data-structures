package main

// Range uses to iterate over data structures
func main() {
	nums := []int{1, 2, 3, 4}

	sum := 0

	for i, num := range nums { // i is index, num is value
		println(i)
		sum += num
	}
	println(sum)
}
