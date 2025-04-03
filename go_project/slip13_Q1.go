// to print sum of all even and odd numbers separately between 1 to 100.

package main

import "fmt"

func main() {
	// Initialize variables to store the sum of even and odd numbers
	var evenSum, oddSum int

	// Loop through numbers from 1 to 100
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			// If the number is even, add it to evenSum
			evenSum += i
		} else {
			// If the number is odd, add it to oddSum
			oddSum += i
		}
	}

	// Print the sum of even and odd numbers
	fmt.Println("Sum of even numbers:", evenSum)
	fmt.Println("Sum of odd numbers:", oddSum)
}
