package main

import (
	"fmt"
	"log"
)

func main() {
	divideByZero()
	printSomeShit()

}

func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide(5, 0))
	//fmt.Println(divide(5, 5))
}

func printSomeShit() {
	fmt.Println("sheeeet...")
}

func divide(a, b int) int {
	return a / b
}

// Recover is a built-in function that regains control of a panicking goroutine.
// Recover is only useful inside deferred functions.
// During normal execution, a call to recover will return nil and have no other effect.
// If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
// it's like a try catch block from java where you don't want to exit after some exception occurs.
