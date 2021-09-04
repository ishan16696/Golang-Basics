package main

import(
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

/*
But sometimes, what we want is that any available services should respond in a desirable time,
if it doesn’t, then default case should get executed.
*/


func service1(c chan string) {
	time.Sleep(3 * time.Second)
	c <- "Hello from service 1"
}

func service2(c chan string) {
	time.Sleep(5 * time.Second)
	c <- "Hello from service 2"
}


func main(){
	fmt.Println("main() started", time.Since(start))

	c1:=make(chan string)
	c2:=make(chan string)

	go service1(c1)
	go service2(c2)


	select {
	case res := <-c1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-c2:
		fmt.Println("Response from service 2", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No response received", time.Since(start))
	}

	fmt.Println("main() stopped", time.Since(start))
	
}