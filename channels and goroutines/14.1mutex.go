package main

import(
	"sync"
	"fmt"
	"time"
)

var mu sync.Mutex
var done bool

func periodic(){
	for{
		fmt.Println("tick ..")
		time.Sleep(1* time.Second)
		mu.Lock()

		if done{
			mu.Unlock()
			return
		}
		mu.Unlock()
		
	}
}

func main(){
	time.Sleep(1* time.Second)
	fmt.Println("main started..")
	go periodic()

	time.Sleep(7*time.Second)

	mu.Lock()
	done =true
	mu.Unlock()

	fmt.Println("cancelled..")
	time.Sleep(3*time.Second)
}