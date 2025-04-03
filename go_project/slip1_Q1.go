//switch case with arimthmatic operations
package main

import ("fmt")

func main()  {
	var num1 int 
	var num2 int 
	var choice int

	fmt.Printf("enter two numbers : ")
	fmt.Scan(&num1,&num2)

	fmt.Printf("enter your choice : ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("addition of 2 numbers is %d\n",num1+num2)
	case 2:
		fmt.Printf("subtraction of 2 numbers is %d\n",num1-num2)
	case 3:
		fmt.Printf("multiplication of 2 numbers is %d\n",num1*num2)
	case 4:
		fmt.Printf("division of 2 numbers is %d\n",num1/num2)
	case 5:
		fmt.Printf("modulus of 2 numbers is %d\n",num1%num2)
	default:
		fmt.Printf("wrong input")
	}
}