package main

import(
	"fmt"
	"sync"
)


func main(){
	var wg sync.WaitGroup 
	var a string
	wg.Add(1)

	go func(){
		a = "can able to mutate the string var define in outside of function through inside anonymous function"
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(a)
}