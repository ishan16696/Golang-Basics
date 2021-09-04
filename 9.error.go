package main

import(
	"fmt"
	"errors"
)

func divide(a int,b int) (int,error){

	if b==0{
		return 0,errors.New("Can't divide by zero")
	}else{
		return a/b,nil
	}
}

func main(){

	if result,err:= divide(8,0); err!=nil{
		fmt.Printf("Error occured %v\n",err)
	}else{
		fmt.Printf("%v\n",result)
	}

}