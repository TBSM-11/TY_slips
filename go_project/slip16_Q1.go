// to create a user defined package to find out the area of a rectangle.

rectangle_area/
│
├── rectangle/
│   └── rectangle.go
│
└── main.go

// rectangle/rectangle.go
package rectangle

// Function to calculate the area of a rectangle
func Area(length, width float64) float64 {
    return length * width
}

// main.go
package main

import (
    "fmt"
    "rectangle"
)

func main() {
    // Define length and width
    length := 5.0
    width := 3.0

    // Call the Area function from the rectangle package
    area := rectangle.Area(length, width)

    // Print the result
    fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
