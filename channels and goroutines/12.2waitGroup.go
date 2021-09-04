package main

import(
	"fmt"
	"sync"
	//"time"
)

func service(wg *sync.WaitGroup,instance int){
	//time.Sleep(2 * time.Millisecond)
	fmt.Println("Service called on instance", instance)
	wg.Done() // decrement counter
}

func sendRPC(i int){
	fmt.Printf("Inside RPC: %d ",i)
}

func main(){
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go service(&wg,i)
		go func(x int){
			sendRPC(x)
			wg.Done()
		}(i)
		
	}
	wg.Wait()

}