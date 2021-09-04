package main

import(	
	"fmt"
	"os"
	"io/ioutil"
	"path"
)

func copyFile(sourceFilePath, destinationFilePath string,filePermission os.FileMode) error {
	data, err := ioutil.ReadFile(sourceFilePath)
	if err != nil {
		return fmt.Errorf("unable to read source file %s: %v", sourceFilePath, err)
	}

	err = ioutil.WriteFile(destinationFilePath, data, filePermission)
	if err != nil {
		return fmt.Errorf("unable to create destination file %s: %v", destinationFilePath, err)
	}
	
	return nil
}

func copyDir(sourceDirPath string,destinationDirPath string) error {
	if len(sourceDirPath)==0 || len(destinationDirPath)==0 {
		return nil 
	}

	files,err:= ioutil.ReadDir(sourceDirPath)
	if err!=nil{
		return err
	}

	for _,file:= range files{
		sourcePath := path.Join(sourceDirPath, file.Name())
                destPath := path.Join(destinationDirPath, file.Name())

		
		fileInfo,err := os.Stat(sourcePath)
		if err!= nil{
			return err
		}
	        
                // only handles dir and files
		if fileInfo.Mode().IsDir(){
			os.Mkdir(destPath,fileInfo.Mode())
			copyDir(sourcePath,destPath)
		}else{
			
			copyFile(sourcePath,destPath,fileInfo.Mode())
		}
	}
	return nil
}



func main(){
	path1:= os.Args[1]
	path2:= os.Args[2]

	copyDir(path1,path2)
}
