package main

import "fmt"

func greet(roc <- chan string){
	data:=<-roc
	fmt.Println("Helloo: ",data)
}

func main(){
	fmt.Println("main started..")
	c:=make(chan string)

	go greet(c)
	c<-"Ishan"

	fmt.Println("main stopped..")

}

/*
Go provide easier syntax to convert bi-directional channel to unidirectional channel.
just like c is bidirectional channel and greet function cconverts it into read-only channel

*/