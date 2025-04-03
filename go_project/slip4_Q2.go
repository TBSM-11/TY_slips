// sort array elements in ascending order

package main

import ("fmt"
		"sort"
	)

func main()  {
	arr := []int{2,5,7,1,9}

	sort.Ints(arr)
	fmt.Println(arr)
}