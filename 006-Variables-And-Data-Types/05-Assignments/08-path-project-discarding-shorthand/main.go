package main

import (
	"fmt"
	"path"
)

func main() {
	// 'project/assets/css/' is discarded.
	_, file := path.Split("project/assets/css/main.css")

	fmt.Println("file:", file)
}
