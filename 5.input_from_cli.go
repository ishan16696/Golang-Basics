package main

import(
	"fmt"
	"os"
	"strconv"
)


func main(){

	a   := os.Args[1]
	b,_ := strconv.Atoi(os.Args[2]) // string to int conversion
	c,_ := strconv.ParseInt(os.Args[2],10,64) // base 10 , up to 64bits

	fmt.Printf("%v %T\n",a,a)
	fmt.Printf("%v %T\n",b,b)
	fmt.Printf("%v %T\n",c,c)

}