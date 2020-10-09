// source -> https://medium.com/rungo/anatomy-of-goroutines-in-go-concurrency-in-go-a4cb9272ff88

package main

import(
	"fmt"
	"time"
)

var start time.Time

func init(){
	start = time.Now()
}

func getchars(s string){
	for _,v:= range s{
		fmt.Printf("%c exec at time %v\n",v,time.Since(start))
		time.Sleep(10*time.Millisecond)
	}
	
}

func getdigits(a[] int){
	for _,d:= range a{
		fmt.Printf("%d exec at time %v\n",d,time.Since(start))
		time.Sleep(15*time.Millisecond)
	}
	
}


func main(){
	fmt.Println("main execution started")

	go getchars("hello")

	a:= []int{4,23,1,10}
	go getdigits(a)

	time.Sleep(200*time.Millisecond)


	fmt.Println("main execution stopped")
}
