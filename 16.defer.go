package main

import "fmt"

// defer  statement is used to defered the func call to the end
// when 2 or more defer statement is there, then it executes in LIFO.
// very useful in unlocking a mutex or closing a file.

func foo(){
	defer bar()
	defer bar2()
	fmt.Println("inside foo..")
}

func bar(){
	fmt.Println("inside bar..")
}

func bar2(){
	fmt.Println("inside bar2..")
}



func main(){
	fmt.Println("main started...")
	foo()
	fmt.Println("main stopped..")
}