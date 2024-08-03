package main

import "net/http"

func handlerError(w http.ResponseWriter, h *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
