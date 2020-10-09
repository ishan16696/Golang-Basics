package main

import(
	"fmt"
)

func main(){
	s:= "this is a string"

	fmt.Printf("%v %v\n", string(s[3]),s[3]) //s[2] give the byte , so we have to convert it into string using string()

	b:= []byte(s)  // convert string into array of bytes

	fmt.Println(b)
	s2:= "hello world.."
	fmt.Println(s+s2) // to concatenate can be done with "+" operator
}