package main

import(
	"fmt"
	"strconv"
)

var a int=23 // to declare the package level variable, a typeof global variable

var I int =42 // its visibility is outside the package --> it is public

func main(){

	const myConst int =42 // constant
	var i int=43

	var j float32
	j= float32(i) // typecasting

	l := string(i) // it will give unicode value to l

	l= strconv.Itoa(i) // to convert int into string
	

	fmt.Printf("%v %T\n",j,j)
	fmt.Printf("%v %T\n",l,l)

}