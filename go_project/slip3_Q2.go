// to accept n records of employee info & display records of emp with highest salary

package main

import "fmt"

type employee struct {
	eno int
	name string
	salary int 
}
func main()  {
	var n int
	fmt.Printf("enter the no. of employees : ")
	fmt.Scan(&n)

	employees := make([]employee, n)
	var maxSalary int
	
	for i := 0; i < n; i++ {
		fmt.Printf("enter details of employee %d (eno, name, salary): ",i+1)
		fmt.Scan(&employees[i].eno,&employees[i].name,&employees[i].salary)
	
	if employees[i].salary > maxSalary {
		maxSalary = employees[i].salary
	}
	}

	fmt.Println("\nEmployees with maximum salary:")
	for _,emp := range employees {
		if emp.salary == maxSalary {
			fmt.Printf("Eno : %d, Name : %s, salary : %d\n",emp.eno,emp.name,emp.salary)
		}
	}
}