package main

import (
	"fmt"
	"slices"
)

// Slices are dyanmic arrays
func main() {
	var nums = make([]int, 0, 10)

	nums = append(nums, 1, 2, 3, 4, 5)

	fmt.Println(cap(nums)) // Capacity of slice
	fmt.Println("Hello, World!", nums)

	var nums2 = []int{1, 2, 3, 4, 5} // another way to initialize
	fmt.Println(len(nums2))          // Length of slice

	// Slice operator
	var nums3 = []int{1, 2, 3, 4, 5}
	fmt.Println("slice operator", nums3[0:3])

	// Check slices are equal or not
	fmt.Println(slices.Equal(nums, nums2)) // Returns boolean

}
