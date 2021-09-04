package main

import(
	"fmt"
)

func getSum(a[] int,n int) (int,error) {
	sum:=0

	for i := 0; i < n; i++ {
		sum =sum +a[i]
	}

	return sum,nil
}


func main(){

	//var grades[4] int

	//grades:= [4]int{33,10,99,12}

	a1:= [5]string{"English","french","hindi","tamil","sanskrit"}

	a2:=a1 // a copy of a1 is assigned to a2
	a2[1]="german"

	fmt.Println("a1: ",a1)
	fmt.Println("a2: ",a2)

	var n int
	fmt.Scan(&n)

	a:= make([]int,n) // making an dynamic array

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var p* int // pointer to an int
	
	var sum int
	p = &sum
	*p,_ = getSum(a,n)
	fmt.Println("Sum of array is: ",*p)

	temp:=[4]float32{1800.99,33.33,77.11,100.77}

	sum=0
	for _ ,value :=range temp {
		sum = sum +int(value)
	}

	fmt.Println("Sum of array is: ",sum)

}

