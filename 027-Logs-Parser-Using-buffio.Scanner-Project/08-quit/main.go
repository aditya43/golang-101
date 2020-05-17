// Quit
//
//  Create a program that quits when a user types the
//  same word twice.
//  The program should work case insensitive.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	words := make(map[string]bool)
	for in.Scan() {
		w := strings.ToLower(in.Text())

		if words[w] {
			fmt.Println("TWICE!")
			return
		}
		words[in.Text()] = true
	}
}
