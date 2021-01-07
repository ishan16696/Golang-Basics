package main

import(
	"fmt"
	"time"
)


func main(){

	timer1 := time.NewTimer(2*time.Second)

	fmt.Println("ishan")
	// You tell the timer how long you want to wait, and it provides a channel that will be notified at that time. 
	// This timer will wait 2 seconds.
	// <-timer1.C blocks on the timer’s channel C until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("timer1 is fired")

	timer2 := time.NewTimer(2*time.Second)
	
	go func(){
		<- timer2.C
		fmt.Println("timer2 is fired")

	}()

	//time.Sleep(2 * time.Second)
	stop2:=timer2.Stop() // can cancel the timer before it fires.

	fmt.Println(stop2)
	if stop2{
		fmt.Println("timer2 is stopped")
	}

	time.Sleep(2 * time.Second)


}