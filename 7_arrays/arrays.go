package main

import "fmt"

func main() {
	var nums [4]int                               // Fixed size array // Default values are 0
	nums[0] = 1                                   // assign value to 0th index
	fmt.Println("length of index is", len(nums))  // get length of array
	fmt.Println("value at 0th index is", nums[0]) // get value of 0th index
	fmt.Println("array is", nums)

	nums2 := [3]int{1, 2, 3} // Fixed size array // Default values are 0

	fmt.Println("array is", nums2)

	nums3 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("array is", nums3)
}
