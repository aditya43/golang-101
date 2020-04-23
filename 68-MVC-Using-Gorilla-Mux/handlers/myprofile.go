package handlers

import "net/http"

func MyProfileHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("profile"))

}
