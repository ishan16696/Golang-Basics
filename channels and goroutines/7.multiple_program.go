package main

import(
	"fmt"
)

func square(c chan int){
	fmt.Println("Inside square func...")
	data:=<-c

	c<- data*data
}


func cube(c chan int){
	fmt.Println("Inside cubic func...")
	data:=<-c

	c<- data*data*data

}

func main(){
	fmt.Println("[main] main started...")
	c1:=make(chan int)
	c2:=make(chan int)

	go square(c1)
	go cube(c2)

	testNum:=3

	fmt.Println("Sending no to square")
	c1<-testNum

	fmt.Println("Resuming.")

	fmt.Println("Sending no to cube")
	c2<-testNum

	fmt.Println("Resuming..")
	fmt.Println("Reading from the channel..")

	squareVal:=<-c1
	cubeVal:=<-c2

	fmt.Printf("%v  %v\n",squareVal,cubeVal)

	fmt.Println("main stopped...")


}