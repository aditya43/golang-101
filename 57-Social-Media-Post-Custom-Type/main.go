package main

import (
	"encoding/json"
	"fmt"

	"github.com/aditya43/golang/57-Social-Media-Post-Custom-Type/socialmedia"
)

func main() {
	post := socialmedia.NewPost(
		"Aditya Hajare",
		socialmedia.Moods["thrilled"],
		"नमस्ते आदित्य",
		"नमस्ते आदित्य! A sample post",
		"http://adiinviter.com",
		"",
		"",
		[]string{"नमस्ते", "आदित्य", "नमस्ते आदित्य a sample post"},
	)

	// fmt.Printf("post: %#v\n", post) // Printing raw output

	// Pretty print by JSON
	res, _ := json.MarshalIndent(post, "", "\t")
	fmt.Print(string(res))
}
