// simple approach to send a cancel signal to another go routine using channel 

package main

import(
	"fmt"
	"time"
	"runtime"
)

var signal = make(chan bool)

func printNumber(){
	counter:=1

	for{
		select{
		case <-signal:
			return // kill the go routine
		default:
			time.Sleep(100*time.Millisecond)
			fmt.Println("tick:",counter)
			counter++
		}
	}

}


func main(){

	fmt.Println("main executed...")

	go printNumber()
	time.Sleep(1*time.Second)

	signal<-true

	fmt.Println("Rountines running",runtime.NumGoroutine())
	fmt.Println("main stopped..")
}



/*
Above methods works fine but only when we are sure 2 go routines are there.
But it could be the case that when one goroutine can start other goroutines,and those goroutines start other goroutines and so on, 
then the first goroutines should be able to send cancellation signals to all spawned goroutines.
The sole purpose of the context package is to carry out the cancellation signal across goroutines no matter how they were spawned.
*/
