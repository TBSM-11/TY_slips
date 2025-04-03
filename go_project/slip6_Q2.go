package main

import "fmt"

func main() {
	// Define the source array (or slice)
	sourceArray := []int{1, 2, 3, 4, 5}

	// Create a new slice with the same length as the source array
	destinationArray := make([]int, len(sourceArray))

	// Copy all elements from sourceArray to destinationArray
	copy(destinationArray, sourceArray)

	// Print the destination array
	fmt.Println("Source Array:", sourceArray)
	fmt.Println("Destination Array:", destinationArray)
}
