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

	// 1st: We can also take the average of the total file length
	//      across platforms. It's about 255.
	//
	// https://en.wikipedia.org/wiki/Comparison_of_file_systems#Limits
	//
	// BTRFS   255 bytes
	// exFAT   255 UTF-16 characters
	// ext2    255 bytes
	// ext3    255 bytes
	// ext3cow 255 bytes
	// ext4    255 bytes
	// FAT32   255 bytes
	// NTFS    255 characters
	// XFS     255 bytes
	//
	// total := len(files) * 256

	// 1st B: To be exact, find the total size of all the empty files
	var total int
	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1 // +1 for the newline character when printing the filename afterward
		}
	}

	fmt.Printf("Total required space: %d bytes.\n", total)

	// 2nd: allocate a large enough byte slice in one go
	fileNames := make([]byte, 0, total)

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
