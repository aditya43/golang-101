package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	rx := regexp.MustCompile(`[^a-z]`)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		word = rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	query := `true`
	if _, ok := words[query]; ok {
		fmt.Println("Word Found!")

		return
	}

	fmt.Println("Word not found!")

	for word := range words {
		fmt.Println(word)
	}
}
