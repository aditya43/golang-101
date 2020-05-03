package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/aditya43/golang/68-MVC-Using-Gorilla-Mux/models/socialmedia"
)

func FetchPostsEndpoint(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement fetching posts for a given user

	// We are going to create some mock data and send it out in JSON
	// format.

	// We will actually implement this endpoint, when we cover database
	// persistence later in the course.

	v := mux.Vars(r)

	if v["username"] == "Aditya Hajare" {

		mockPosts := make([]socialmedia.Post, 3)

		post1 := socialmedia.NewPost(
			"Aditya Hajare",
			socialmedia.Moods["thrilled"],
			"नमस्ते आदित्य",
			"नमस्ते आदित्य! A sample post",
			"http://adiinviter.com",
			"",
			"",
			[]string{"नमस्ते", "आदित्य", "नमस्ते आदित्य a sample post"},
		)

		post2 := socialmedia.NewPost(
			"Aditya Hajare",
			socialmedia.Moods["happy"],
			"नमस्ते आदित्य",
			"नमस्ते आदित्य! A sample post",
			"http://adiinviter.com",
			"",
			"",
			[]string{"नमस्ते", "आदित्य", "नमस्ते आदित्य a sample post"},
		)

		post3 := socialmedia.NewPost(
			"Aditya Hajare",
			socialmedia.Moods["hopeful"],
			"नमस्ते आदित्य",
			"नमस्ते आदित्य! A sample post",
			"http://adiinviter.com",
			"",
			"",
			[]string{"नमस्ते", "आदित्य", "नमस्ते आदित्य a sample post"},
		)

		mockPosts = append(mockPosts, *post1)
		mockPosts = append(mockPosts, *post2)
		mockPosts = append(mockPosts, *post3)
		json.NewEncoder(w).Encode(mockPosts)

	} else {
		json.NewEncoder(w).Encode(nil)

	}
}
