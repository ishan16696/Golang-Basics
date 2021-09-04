// callback is when we pass a func into a func as an argument in golang.

package main

import (
	"fmt"
)

func main() {

	g := func(s string) int {
		if len(s) == 0 {
			return 0
		}
		if len(s) == 1 {
			return 1
		}
		return len(s)
	}

	x := foo(g, "Ishan")
	fmt.Println(x)
}

func foo(f func(s string) int, s string) int {
	n := f(s)
	n++
	return n
}
