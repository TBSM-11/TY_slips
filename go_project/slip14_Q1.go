// to demonstrate working of slices (like append, remove, copy etc.)

package main

import (
	"fmt"
)

func main() {
	// Initialize a slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial slice:", slice)

	// Append elements to the slice
	slice = append(slice, 6, 7, 8)
	fmt.Println("After appending 6, 7, 8:", slice)

	// Remove an element from the slice (let's remove the element at index 2)
	// Removing element at index 2 means we want to exclude the number 3.
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("After removing the element at index 2:", slice)

	// Copy elements from one slice to another
	// Create a new slice with the same length as the original slice
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Println("Copied slice:", newSlice)

	// Modify the original slice to demonstrate that the copied slice is independent
	slice[0] = 100
	fmt.Println("Modified original slice:", slice)
	fmt.Println("Unmodified copied slice:", newSlice)

	// Resize slice (this example shrinks the slice to the first 3 elements)
	slice = slice[:3]
	fmt.Println("After resizing to first 3 elements:", slice)
}
