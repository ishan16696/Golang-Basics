package main

import(
	"os"
	"fmt"
	"io/ioutil"
)

func listFiles(sourcePath string) error{
	if len(sourcePath)==0 {
		fmt.Println("Source Path can't be null")
		return nil
	}

	files, _ := ioutil.ReadDir(sourcePath)

	for _, file := range files {
		fmt.Printf("%v %vk  %v\n",file.Mode(),file.Size()/1000,file.Name())
	}

	return nil
}

func main() {

	
	//dirPath := path.Join("./Desktop",os.Args[1]) 
	dirPath := os.Args[1]
	listFiles(dirPath)

}