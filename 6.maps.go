package main

import(
	"fmt"
)

func main(){

	score := make(map[string]int)

	score["ishan"]=69
	score["tyagi"]=96
	score["ayush"]=77


	fmt.Println(score)

	fmt.Println("Ishan score is:" ,score["ishan"])
	
	delete(score,"ishan") // to delete the key
	for k,v := range score{
		fmt.Println(k,v)
	}

}