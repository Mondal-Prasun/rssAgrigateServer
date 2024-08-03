package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, h *http.Request) {
	respondWithJson(w, 200, struct{}{})
}
