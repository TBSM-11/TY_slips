package main

import "fmt"

// Define the Student structure
type Student struct {
	Name    string
	RollNo  int
	Marks   float64
}

// Define the Show method with a pointer receiver
func (s *Student) Show() {
	// Display the student's details
	fmt.Printf("Student Name: %s\n", s.Name)
	fmt.Printf("Roll Number: %d\n", s.RollNo)
	fmt.Printf("Marks: %.2f\n", s.Marks)
}

func main() {
	// Create a Student instance
	student1 := &Student{
		Name:   "John Doe",
		RollNo: 101,
		Marks:  89.5,
	}

	// Call the Show method using the pointer
	student1.Show()
}
