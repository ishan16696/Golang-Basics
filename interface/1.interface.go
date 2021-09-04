package main

import (
	"fmt"
)

type circle struct{
	radius float64
}

type rectangle struct{
	height float64
	width float64
}

func (c circle) area() float64{
	return c.radius*3.14
}


func (r rectangle) area() float64{
	return r.height*r.width
}

// a value can be more than 1 type
type shape interface{
	area() float64
}

func abs(s shape) {
	fmt.Println(s.area())
}


func NewCircle(radius float64) shape{
 return circle{
 	radius: radius,
 }
}

func main() {
	c:= circle{
		radius: 9.1,
	}
	r:=rectangle{
		height: 5.1,
		width: 11.9,
	}
	
	abs(c)
	abs(r)

	s:= NewCircle(7.2)
	fmt.Println(s.area())
}
