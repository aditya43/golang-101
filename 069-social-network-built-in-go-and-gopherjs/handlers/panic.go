package handlers

import (
	"net/http"
)

func TriggerPanicHandler(w http.ResponseWriter, r *http.Request) {

	panic("Triggering a Panic!")
}
