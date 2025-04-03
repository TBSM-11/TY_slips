// to illustrate the concept of
// returning multiple values from a function. ( Add, Subtract,Multiply, Divide)

package main

import (
	"fmt"
	"log"
)

// Function that returns multiple values for addition, subtraction, multiplication, and division
func calculate(a, b float64) (float64, float64, float64, float64, error) {
	// Perform the operations
	add := a + b
	subtract := a - b
	multiply := a * b
	// Check for division by zero
	if b == 0 {
		return add, subtract, multiply, 0, fmt.Errorf("cannot divide by zero")
	}
	// Perform division
	divide := a / b
	return add, subtract, multiply, divide, nil
}

func main() {
	// Example values for a and b
	a := 10.0
	b := 5.0

	// Call the calculate function
	add, subtract, multiply, divide, err := calculate(a, b)

	// Check for error in division
	if err != nil {
		log.Fatal(err)
	}

	// Print the results of all operations
	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, add)
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, subtract)
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, multiply)
	fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, divide)
}
