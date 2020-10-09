package main

import(
	"fmt"
)


type User struct{
	name string
	id string
	age int
}

func main(){

	candidates:= []User{{"Ishan","99ef",23},{"ayush","99if",28}}

	fmt.Println(candidates)
	rahul:= User{"Rahul","rrk",77}
	fmt.Println(rahul)

	rob := new(User)
	rob.name= "Rob"
	rob.id= "iou"
	rob.age= 19

	fmt.Println(*rob)
	fmt.Println((*rob).name)

}