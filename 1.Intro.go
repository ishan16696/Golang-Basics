package main

import "fmt"


func main(){
	
	//<keyword> <identifier> <type>
	var name string
	var age int

	rollNo := 99

	fmt.Scan(&name)
	fmt.Println("you entered the name: ",name)

	fmt.Scan(&age)
	
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age>12 && age <18{
		fmt.Println("You are an teenager")
	}else{
		fmt.Println("You are not an teenager")
	}

	fmt.Printf("Roll no is: %d\n", rollNo)


}
