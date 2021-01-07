package main

import(
	"fmt"
)

func foo(x int){
	fmt.Println("foo..",x)
}

//function taking argument as function
func test2(myfunc func(int) int){

	// calling the function which we received as argument
	fmt.Println(myfunc(7))
}



func main(){
	x:=foo
	x(10) // it is same as foo(10)

	test:=func (x int) int{
		return x*-1
	}(2) // () -> this is used to call the function
	fmt.Println(test)

	temp:= func(x int) int{
		return x*2
	}

	// passing function as argument to other function
	test2(temp) 
}