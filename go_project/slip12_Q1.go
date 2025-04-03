// Write a program in GO language to swap two numbers using call by reference concept

package main

import "fmt"

// Function to swap two numbers using pointers (call by reference)
func swap(a *int, b *int) {
	// Temporary variable to hold the value of 'a'
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	// Declare two numbers
	var num1, num2 int

	// Accept the values from the user
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Display original values
	fmt.Println("Before swap:")
	fmt.Println("num1 =", num1, "num2 =", num2)

	// Call the swap function with pointers
	swap(&num1, &num2)

	// Display swapped values
	fmt.Println("After swap:")
	fmt.Println("num1 =", num1, "num2 =", num2)
}
