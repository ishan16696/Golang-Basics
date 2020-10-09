package main

import(
	"fmt"
)

func squares(c chan int){

	for i := 0; i < 10; i++ {
		c<- i*i
	}

	close(c)

}

func main(){
	fmt.Println("main started...")
	c:=make(chan int)

	go squares(c)

	for {
		val,ok:= <-c    // you can also use  val,ok:= range c
		if ok==false{
			fmt.Println("Loop break ",val,ok)
			break
		}else{
			fmt.Println(val,ok)
		}
	}

	fmt.Println("main stopped....")
}