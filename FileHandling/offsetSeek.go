package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		fmt.Errorf("unable to open the file %v", err)
	}

	var currentOffset int64
	if currentOffset, err = file.Seek(0, io.SeekCurrent); err != nil {
		fmt.Errorf("unable to set the offset for file %v", err)
	}
	fmt.Printf("Starting offset of file is: %d\n", currentOffset)

	if currentOffset, err = file.Seek(0, io.SeekEnd); err != nil {
		fmt.Errorf("unable to set the offset for file %v", err)
	}
	fmt.Printf("ending offset of file is: %d\n", currentOffset)

	// Seek sets the file pointer for the next Read or Write on file to given offset
	// example: here file pointer is set to read after 5th byte from starting
	if currentOffset, err = file.Seek(5, io.SeekStart); err != nil {
		fmt.Errorf("unable to set the offset for file %v", err)
	} else {
		fmt.Printf("currentOffset of file is: %d\n", currentOffset)
	}

	// read next 13 byte of data
	data := make([]byte, 13)
	if _, err := file.Read(data); err != nil {
		fmt.Errorf("unable to read the file %v", err)
	}
	fmt.Printf("Content with file seek: %s\n", string(data))

	//var currentOffset int64
	if currentOffset, err = file.Seek(0, io.SeekCurrent); err != nil {
		fmt.Errorf("unable to set the offset for file %v", err)
	}
	fmt.Printf("currentOffset of file is: %d\n", currentOffset)

	// set the file pointer to before 7th byte from end.
	if offset, err := file.Seek(-7, io.SeekEnd); err != nil {
		fmt.Errorf("unable to set the offset for file %v", err)
	} else {
		fmt.Printf("currentOffset of file is: %d\n", offset)
	}

	// read next 10 byte of data
	data = make([]byte, 10)
	if _, err := file.Read(data); err != nil {
		fmt.Errorf("unable to read the file %v", err)
	}
	fmt.Printf("Content with of file from end: %s\n", string(data))
}
