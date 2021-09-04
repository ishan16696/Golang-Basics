package main

//A value can be more than one type


import(
	"fmt"
	"math"
)

// <keyword> <identifier> <type>
type shapes interface{
	area() float64
}

type circle struct{
	radius float64
}

type rectangle struct{
	height float64
	width float64
}

func (c circle) area() (float64){
	return c.radius*c.radius*math.Pi
}


func (r rectangle) area() (float64){
	return r.height*r.width
}

func main(){
	
	s:= []shapes{
		circle{
			radius: 7.12,
		},
		rectangle{
			height: 8,
			width: 10,
		},
	}

	fmt.Println(s[0].area())
	fmt.Println(s[1].area())

}