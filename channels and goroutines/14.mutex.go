package main

import (
	"fmt"
	"sync"
)

var i int 

// goroutine increment global variable i
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() // acquire lock
	i = i + 1
	m.Unlock() // release lock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	// wait until all 1000 gorutines are done
	wg.Wait()

	// value of i should be 1000
	fmt.Println("value of i after 1000 operations is", i)
}

/*
If we remove the mutex lock then race codition will get occur.because even if goroutine have different stack but still 
have the  goroutine is using the same heap(just like threads)

*/