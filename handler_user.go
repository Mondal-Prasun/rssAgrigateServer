package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mondal-Prasun/rssAgrigateServer/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, h *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(h.Body)
	dataParam := parameters{}

	err := decoder.Decode(&dataParam)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	user, err := apiCfg.db.CreateUser(h.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Createdat: time.Now().UTC(),
		Updatedat: time.Now().UTC(),
		Name:      dataParam.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
	}

	respondWithJson(w, 200, user)
}
