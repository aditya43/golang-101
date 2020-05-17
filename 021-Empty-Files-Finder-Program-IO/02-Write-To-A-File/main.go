package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var dirPath string
	var fileNames []byte

	fmt.Println("Please enter directory path:")
	fmt.Scan(&dirPath)

	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.Size() == 0 {
			fileNames = append(fileNames, file.Name()...)
			fileNames = append(fileNames, '\n')
		}
	}

	err = ioutil.WriteFile("out.txt", fileNames, 0644)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
