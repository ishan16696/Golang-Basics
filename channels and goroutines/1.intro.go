package main

// channel is declared using a keyword "chan"

import(
	"fmt"
)

func main(){
	// var c chan int 	// default value is nil that's why "make" is used 

	c:=make(chan int)
	fmt.Println(c)
}

/*
Looks like it is a memory address. Channels by default are pointers. 
Mostly,when you want to communicate with a goroutine,you pass the channel as an argument to the function or method. 
Hence when goroutine receives that channel as an argument, you don’t need to dereference it to push or pull data from that channel.


*/


/*
source:

https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb

*/