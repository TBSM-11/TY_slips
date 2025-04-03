// Write a program in GO language to accept the book details such
// as BookID, Title, Author, Price. Read and display the details of
// ‘n’ number of books

package main

import "fmt"

// Define the Book structure
type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	// Accept the number of books from the user
	var n int
	fmt.Print("Enter the number of books: ")
	fmt.Scan(&n)

	// Create a slice to hold n books
	books := make([]Book, n)

	// Accept the details for each book
	for i := 0; i < n; i++ {
		fmt.Printf("\nEnter details for Book %d:\n", i+1)
		fmt.Print("Enter BookID: ")
		fmt.Scan(&books[i].BookID)
		fmt.Print("Enter Title: ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Enter Author: ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Enter Price: ")
		fmt.Scan(&books[i].Price)
	}

	// Display the details of all books
	fmt.Println("\nBook Details:")
	for i := 0; i < n; i++ {
		fmt.Printf("\nBook %d\n", i+1)
		fmt.Printf("BookID: %d\n", books[i].BookID)
		fmt.Printf("Title: %s\n", books[i].Title)
		fmt.Printf("Author: %s\n", books[i].Author)
		fmt.Printf("Price: %.2f\n", books[i].Price)
	}
}
