// source -> https://medium.com/rungo/anatomy-of-goroutines-in-go-concurrency-in-go-a4cb9272ff88

package main

import(
	"fmt"
	"time"
)

func getchars(s string){
	for _,v:= range s{
		fmt.Printf("%c ",v)
	}
	fmt.Printf("\n")
}

func getdigits(a[] int){
	for _,d:= range a{
		fmt.Printf("%d ",d)
	}
	fmt.Printf("\n")
}


func main(){
	fmt.Println("main execution started")

	go getchars("hello")

	a:= []int{4,23,1,10}
	go getdigits(a)

	time.Sleep(2*time.Millisecond)


	fmt.Println("main execution stopped")
}
