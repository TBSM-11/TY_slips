// to print a multiplication table of number using function

package main

import "fmt"

// Function to print the multiplication table of a number
func printMultiplicationTable(number int) {
	// Loop from 1 to 10 to print the multiplication table
	for i := 1; i <= 10; i++ {
		// Print the result of number * i
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}

func main() {
	// Ask the user for a number
	var number int
	fmt.Print("Enter a number to print its multiplication table: ")
	fmt.Scan(&number)

	// Call the function to print the multiplication table
	printMultiplicationTable(number)
}
