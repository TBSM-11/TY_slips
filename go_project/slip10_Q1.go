// to read and write Fibonacci series to the using channel.

package main

import (
	"fmt"
)

// Function to generate Fibonacci series and send to channel
func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a // Send Fibonacci number to channel
		a, b = b, a+b // Update values of a and b
	}
	close(ch) // Close the channel once done sending all numbers
}

func main() {
	// Define the number of Fibonacci numbers to generate
	var n int
	fmt.Print("Enter the number of Fibonacci numbers to generate: ")
	fmt.Scan(&n)

	// Create a channel to send and receive Fibonacci numbers
	ch := make(chan int)

	// Start a goroutine to generate Fibonacci numbers
	go fibonacci(n, ch)

	// Read and print the Fibonacci numbers from the channel
	fmt.Println("Fibonacci Series:")
	for num := range ch {
		fmt.Println(num)
	}
}
