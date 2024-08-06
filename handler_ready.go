package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, h *http.Request) {

	respondWithJson(w, 200, struct {
		Test string `json:"test"`
	}{
		Test: "Hello I am ready",
	})
}
