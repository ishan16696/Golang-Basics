package main

import(
	"fmt"
)

type name string

type User struct{
	first name
	last name
	id string
	age int
}

func main(){

	candidates:= []User{
		{
			first: "Ishan",
			last: "tyagi",
			id: "is23",
			age: 22,
	},
		{
			first: "ayush",
			last: "tyagi",
			age: 28,
			id: "ty23",
		},
	}

	fmt.Println(candidates)
	rahul:= User{"Rahul","sharma","rrk",77}
	fmt.Println(rahul)

	rob := new(User)
	rob.first= "Rob"
	rob.last= "Chaudhary"
	rob.id= "iou"
	rob.age= 19

	fmt.Println(*rob)
	fmt.Println((*rob).first)

	//pointer to structure
	ishan:= &User{
		first: "Ishan",
		last: "Tyagi",
		id: "07is",
		age: 24,
	}
	fmt.Println(ishan)

}