package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	var dirPath string

	fmt.Println("Please enter directory path:")
	fmt.Scan(&dirPath)

	files, err := ioutil.ReadDir(dirPath)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		if file.Size() == 0 {
			fmt.Println(file.Name())
		}
	}
}
