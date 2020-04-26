package handlers

import "net/http"

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("profile"))
}
