package main

import(
	"fmt"
	"os"
	"log"
)


func main(){

	// create the file os.Args[1] if doesn't exist and open it for reading and writing
	// append the content to the file while writing instead of overriding
	file,err:= os.OpenFile(os.Args[1],os.O_RDWR | os.O_CREATE | os.O_APPEND,0777)

	if err != nil {
		log.Printf("Error..")
	}

	fmt.Fprint(file,os.Args[2])
	fmt.Fprint(file,os.Args[3])
}