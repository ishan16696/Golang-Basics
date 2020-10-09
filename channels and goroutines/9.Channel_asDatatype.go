package main

import(
	"fmt"
)

func greet(c chan string){
	data:=<-c
	fmt.Println("hello: ",data)
}

func greeter(cc chan chan string){
	c:=make(chan string)
	cc <- c

}

func main(){
	fmt.Println("main started..")

	// 'cc' is a channel of dataType channel of string
	cc:=make(chan chan string)

	go greeter(cc)

	c:= <-cc
	go greet(c)

	c<-"Ishan"

	fmt.Println("main stopped..")
	


}