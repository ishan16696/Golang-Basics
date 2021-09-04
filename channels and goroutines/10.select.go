package main

import(
	"fmt"
	"time"
)

var start time.Time

func init(){
	start=time.Now()
}

func service1(c chan string){
	time.Sleep(3*time.Second)
	c <-"Hello from service1"
}

func service2(c chan string){
	time.Sleep(5*time.Second)
	c<-"Hello from service2"
}

func main(){
	fmt.Println("main started.")
	c1:=make(chan string)
	c2:=make(chan string)

	go service1(c1)
	go service2(c2)

	select{
	case res := <-c1:
		fmt.Println("Response from service 1", res, time.Since(start))
	case res := <-c2:
		fmt.Println("Response from service 2", res, time.Since(start))
	// default:
	// 	fmt.Println("No Response from services")
	}

	fmt.Println("main stopped..")
}