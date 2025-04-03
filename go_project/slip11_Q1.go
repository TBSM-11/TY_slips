// to check whether the accepted number is two digit or not.
package main

import "fmt"

// Function to check if the number is a two-digit number
func isTwoDigit(num int) bool {
	// Check if the number is between 10 and 99 or -10 and -99
	return (num >= 10 && num <= 99) || (num <= -10 && num >= -99)
}

func main() {
	// Accept a number from the user
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check if the number is a two-digit number
	if isTwoDigit(num) {
		fmt.Println(num, "is a two-digit number.")
	} else {
		fmt.Println(num, "is not a two-digit number.")
	}
}
