package main

import "fmt"

/*
When an interface has zero methods, it is called an empty interface. 
This is represented by interface{}. Since the empty interface has zero methods, all types implement this interface implicitly.

*/



type MyString string

type Rect struct {
	width  float64
	height float64
}

func explain(i interface{}) {
	fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func main() {
	ms := MyString("Hello World!")
	r := Rect{5.5, 4.5}
	explain(ms)
	explain(r)

	m:=make(map[string]interface{})

	m["ishan"]=8
	m["tyagi"]="Dev"
	m["ilu"]=9.9

	fmt.Printf("%v",m)
}
