package main

import (
	"fmt"
)


func greet(c chan string){
	data:= <-c
	fmt.Println("Hello " + data + " golang") // reading data from channel
}

func main(){
	fmt.Println("main Started...")
	c:=make(chan string)

	go greet(c)

	c<-"Ishan"  // writing data to channel
	
	/*
	When some data is written to the channel,goroutine is blocked until some other goroutine reads it from that channel. 
	*/

	fmt.Println("main Stopped...")
}

/*
the process has 2 goroutines while active goroutine is main goroutine till line 4
We pushed a string value "ishan" to channel c. 
At this point,main goroutine is blocked until some goroutine reads it. 
Go scheduler schedule greet goroutine.
Then main goroutine becomes active and execute the final statement,printing main stopped
*/