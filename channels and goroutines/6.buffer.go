package main

import(
	"fmt"
)

func sum(c chan int){
	sum:=0

	for i := 0; i < 3; i++ {
		data:=<-c
		sum+=data
	}

	fmt.Println(sum)
}

func main(){
	fmt.Println("main started...")
	c:=make(chan int,3)

	go sum(c)
	c<-15 // len=1 cap=3
	c<-21//len=2 cap=3
	fmt.Printf("%v  %v\n",len(c),cap(c)) // length of channel and capacity of channel

	c<-33 // len=3 cap=3
	c<-44 // main goroutine blocks here


	fmt.Println("main stopped...")
}