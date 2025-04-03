//fibonacci series

package main

import "fmt"

func main()  {
	var nTerm int

	fmt.Printf("enter the no. of terms : ")
	fmt.Scan(&nTerm)

	var (
		a int = 0
		b int = 1
	)

	fmt.Printf("Fibonacci series : ")
	
	for i := 0; i < nTerm; i++ {

		fmt.Print(a," ")
		a,b=b,a+b
	}
}