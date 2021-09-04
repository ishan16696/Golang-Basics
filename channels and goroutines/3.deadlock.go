package main

import "fmt"

func main(){
	fmt.Println("main started..")

	c:=make(chan string)

	c<-"Ishan" //if you are to send data to channel,it will block current goroutine and unblock others until some goroutine reads the data from it

	fmt.Println("main stopped...")
}


/*
As discussed,when we write or read data from a channel,that goroutine is blocked and control is passed to available goroutines. 
What if there are no other goroutines available, imagine all of them are sleeping. 
That’s where deadlock error occurs crashing the whole program.
*/