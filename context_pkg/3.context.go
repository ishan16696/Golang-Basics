package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// start time
var startTime = time.Now()


func worker(ctx context.Context, sec int) {
	select {

	
	case <-ctx.Done():
		fmt.Printf("%0.2fs - worker(%ds) killed!\n", time.Since(startTime).Seconds(), sec)
		return // kills goroutine

	
	case <-time.After(time.Duration(sec) * time.Second):
		fmt.Printf("%0.2fs - worker(%ds) completed the job.\n", time.Since(startTime).Seconds(), sec)
	}
}


func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(3*time.Second))    // internally it called context.Deadline() 
	
	defer cancel()

	go worker(ctx, 2) 
	go worker(ctx, 1) 
	go worker(ctx, 8) 
	
	time.Sleep(5 * time.Second)
	
	fmt.Println("Number of active goroutines", runtime.NumGoroutine())
}
