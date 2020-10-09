package main

import(
	"fmt"
)

type Point struct{
	X float32
	Y float32
}

//function
func isAbove(p Point, y float32) bool{
	return p.Y > y
}

//method with receiver 'Point'
func (p Point) isAbove(y float32) bool{
	return p.Y > y
}

func main(){
	
	p:= Point{2.1,4.2}

	q:= new(Point)
	q.X=2.7
	q.Y=0.8

	fmt.Println("Point: ",p)
	fmt.Println("Is point p located above the line y=1.0 :", p.isAbove(1))


	fmt.Println("Is point p located above the line y=1.0 :", isAbove(p,1))


	fmt.Println("Point: ",*q)
	fmt.Println("Is point q located above the line y=1.0 :", (*q).isAbove(1))


	fmt.Println("Is point q located above the line y=1.0 :", isAbove(*q,1))
}

