package main

import(
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup,instance int){
	time.Sleep(2 * time.Second)
	fmt.Println("Service called on instance", instance)
	wg.Done() // decrement counter
}

func main(){

	fmt.Println("main started..")
	var wg sync.WaitGroup 

	for i := 0; i < 3; i++ {
		wg.Add(1)  // increment counter
		go service(&wg,i)
	}

	wg.Wait() // blocks here

	fmt.Println("main stopped..")
	
}

/*

WaitGroup is struct with counter value which tracks how many goroutines were spawned and how many have completed their job.
This counter when reaches zero, means all goroutines have done their job.

it has 3 methods: 
1) Add 
2)Wait 
3)Done

1)Add method expects one int argument which is delta for the WaitGroup counter. 
The counter is nothing but an integer with default value 0. It holds how many goroutines are running.
When WaitGroup is created, its counter value is 0 and we can increment it by passing delta as parameter using Add method.
Remember, counter does not increment understand when a goroutine was launched, hence we need to manually increment it.

2)Wait method is used to block the current goroutine from where it was called.Once counter reaches 0,that goroutine will unblock.Hence, we need something to decrement the counter.

3)Done method decrements the counter. It does not accept any argument, hence it only decrements the counter by 1.


After for loop has done executing,still did not pass control to other goroutines.This is done by calling Wait method on wg like wg.Wait(). 
This will block the main goroutine until the counter reaches 0 . Once the counter reaches 0 because from 3 goroutines,
we called Done method on wg 3 times, main goroutine will unblock and starts executing further code.

*/
