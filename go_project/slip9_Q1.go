// palindrome check
package main

import "fmt"

// Function to check if a number is a palindrome
func isPalindrome(n int) bool {
	original := n
	reversed := 0

	// Reverse the number
	for n != 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n = n / 10
	}

	// Check if the original number is equal to the reversed number
	return original == reversed
}

func main() {
	// Accept a number from the user
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	// Check if the number is a palindrome
	if isPalindrome(number) {
		fmt.Println(number, "is a palindrome.")
	} else {
		fmt.Println(number, "is not a palindrome.")
	}
}
