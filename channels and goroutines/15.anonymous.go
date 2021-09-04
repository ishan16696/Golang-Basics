package main

import "fmt"

func main(){
	fmt.Println("main started..")
	c:=make(chan string)

	go func(c chan string){
		data:=<-c
		fmt.Println("Hello ",data)
	}(c)

	c<-"Ishan"

	fmt.Println("main stopped...")
}