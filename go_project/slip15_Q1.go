// to demonstrate function return multiple values.

package main

import "fmt"

// A function that returns multiple values
func addAndSubtract(a, b int) (int, int) {
    sum := a + b
    difference := a - b
    return sum, difference
}

func main() {
    // Calling the function with two arguments
    sum, diff := addAndSubtract(10, 5)

    // Printing the results
    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", diff)
}
