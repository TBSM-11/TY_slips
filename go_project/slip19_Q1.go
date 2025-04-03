// to illustrate the function returning multiple values(add, subtract).

package main

import "fmt"

// Function to perform both addition and subtraction
func addSubtract(a, b int) (int, int) {
    sum := a + b
    difference := a - b
    return sum, difference
}

func main() {
    // Define two integers
    num1 := 10
    num2 := 5

    // Call the function to get both sum and difference
    sum, difference := addSubtract(num1, num2)

    // Print the results
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
}
