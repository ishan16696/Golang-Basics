package main

import(
	"fmt"
)

func greet(c chan string){
	
	data1 := <-c
	fmt.Println("Read data from channel ",data1)
	data2 := <-c
	fmt.Println("Read 2nd data from channel ",data2)
	 
}


func main(){
	fmt.Println("main started...")
	var data string 
	fmt.Scan(&data)

	c := make(chan string)

	go greet(c)

	c<-data

	//close(c) // closing a channel --> it will cause an error
	fmt.Println("After first send")

	c <-"ishan"

	fmt.Println("main stopped...")

}

/*

To help you understand blocking concept,first send operation c <-data is blocking and some goroutine has to read data from the channel, 
hence greet goroutine is scheduled by the Go Scheduler.
Then first read operation <-c is non-blocking because data is present in channel c to be read from.
Second read operation <-c will be blocking because channel c does not have any data to be read from,hence Go Scheduler
activates main goroutine and program starts execution from close(c) function

*/