package main

import(
	"fmt"
)

type Point struct{
	X float32
	Y float32
}

func (p* Point) Translate(dx,dy float32){
	p.X= p.X + dx
	p.Y = p.Y +dy
}


func main(){
	
	q:= Point{3,4}

	fmt.Println("Point is: ", q)
	q.Translate(7,9)

	fmt.Println("After Translate: ",q)

}