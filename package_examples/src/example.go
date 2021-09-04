package main

import(
	"fmt"
	"package_examples/src/numbers"
	"package_examples/src/stringUtils" 
)

func main(){
	fmt.Println(numbers.IsPrime(19))
	fullname,length := stringUtils.Fullname("Ishan","Tyagi")
	fmt.Println(fullname,length)
}

