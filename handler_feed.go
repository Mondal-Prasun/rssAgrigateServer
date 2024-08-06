package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mondal-Prasun/rssAgrigateServer/internal/database"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, h *http.Request, theUser database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(h.Body)
	dataParam := parameters{}

	err := decoder.Decode(&dataParam)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	feed, err := apiCfg.db.CreateFeed(h.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Createdat: time.Now().Local(),
		Updatedat: time.Now().Local(),
		Name:      dataParam.Name,
		Url:       dataParam.Url,
		Userid:    theUser.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
	}

	parsedFeed := Feed{
		ID:        feed.ID,
		Createdat: feed.Createdat,
		Updatedat: feed.Updatedat,
		Name:      feed.Name,
		Url:       feed.Url,
		Userid:    feed.Userid,
	}

	respondWithJson(w, 201, parsedFeed)
}

func (apiCfg *apiConfig) handlerGetAllFeed(w http.ResponseWriter, h *http.Request) {

	feeds, err := apiCfg.db.GetFeeds(h.Context())
	if err != nil {
		respondWithError(w, 404, fmt.Sprintf("Cannot fetch feeds: %v", err))
	}

	alteredFeeds := []Feed{}

	for _, value := range feeds {
		alteredFeeds = append(alteredFeeds, Feed{
			ID:        value.ID,
			Createdat: value.Createdat,
			Updatedat: value.Updatedat,
			Name:      value.Name,
			Url:       value.Url,
			Userid:    value.Userid,
		})
	}

	respondWithJson(w, 201, alteredFeeds)
}
