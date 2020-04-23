package handlers

import (
	"net/http"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {
	fooID := r.Context().Value("fooID").(string)
	_ = w.Write([]byte(fooID))
}
