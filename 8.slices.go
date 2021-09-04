package main

import(
	"fmt"
	"sort"
)

func main(){

	myList:= []int{18000,2750,13000,2100}

	myList = append(myList,99)
	sort.Ints(myList)
	fmt.Println(myList)

	var n int
	fmt.Scan(&n)
	var mySlices[] int
	var x int
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		mySlices = append(mySlices,x)
	}

	fmt.Println(mySlices)

}