package main

import(
	"fmt"
	"time"
	"context"
)

var c = make(chan int)

func square(ctx context.Context){
	i:=0
	for{
		select{
		case <-ctx.Done():
			return // kill the go routine
		default:
			c<-i*i
			i++
		}
	}
}

func main(){

	fmt.Println("main started...")

	ctx ,cancel := context.WithCancel(context.Background())

	go square(ctx)

	for i := 0; i < 5; i++ {
		fmt.Println("Next square is", <-c)
	}

	// cancel context
	cancel() 

	
	time.Sleep(3 * time.Second)

	fmt.Println("main stopped...")

}


// context.Background() and context.TODO() is just empty context used to pass in place of parent context.
// context will get cancel when cancel() will be called or parent context will get cancel.
// when cancel() is called, it will send a signal to context.Done() channel.