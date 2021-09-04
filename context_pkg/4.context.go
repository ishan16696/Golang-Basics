package main

import(
	"fmt"
	"time"
	"context"
)

var c = make(chan int)

func cubic(ctx context.Context){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("Cubic shutting down ...")
			return
		default:
			time.Sleep(100*time.Millisecond)
			fmt.Println("inside cubic..")
		}
	}
}


func square(ctx context.Context){
	for{
		select{
		case <-ctx.Done():
			fmt.Println("square shutting down ...")
			return // kill the go routine
		default:
			time.Sleep(1*time.Second)
			go cubic(ctx)
			fmt.Println("inside square..")
			
		}
	}
}

func cancelAllThreads(cancel context.CancelFunc){
	fmt.Println("Cancelling the threads...")
	cancel()
}

func main(){

	fmt.Println("main started...")

	ctx ,cancel := context.WithCancel(context.Background())
	
	go square(ctx)

	for i := 0; i < 5; i++ {
		fmt.Println("Next square is:",i*i )
	}

	time.Sleep(2*time.Second)
	cancelAllThreads(cancel) 
	
	time.Sleep(2*time.Second)

	fmt.Println("main stopped...")

}


// context.Background() and context.TODO() is just empty context used to pass in place of parent context.
// context will get cancel when cancel() will be called or parent context will get cancel.
// when cancel() is called, it will send a signal to context.Done() channel.