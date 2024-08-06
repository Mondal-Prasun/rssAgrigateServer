package main

import (
	"fmt"
	"net/http"

	"github.com/Mondal-Prasun/rssAgrigateServer/internal/database"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg apiConfig) middlewareAuth(authHandler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, h *http.Request) {
		apiKey, err := GetApiKeyFromHeader(h.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Auth error: %v", err))
		}

		theUser, errr := apiCfg.db.GetUserByApikey(h.Context(), apiKey)

		if errr != nil {
			respondWithError(w, 404, fmt.Sprintf("User not found: %v", errr))
		}

		authHandler(w, h, theUser)
	}
}
