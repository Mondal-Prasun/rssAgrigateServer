package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Mondal-Prasun/rssAgrigateServer/internal/database"
	"github.com/go-chi/chi/v5"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, h *http.Request, theUser database.User) {

	type parameters struct {
		FeedId uuid.UUID `json:"feedId"`
	}

	decoder := json.NewDecoder(h.Body)
	dataParam := parameters{}

	err := decoder.Decode(&dataParam)
	if err != nil {
		respondWithError(w, 400, err.Error())
		return
	}
	feedFollow, err := apiCfg.db.CreateFeedFollow(h.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		Updatedat: time.Now().Local(),
		Createdat: time.Now().Local(),
		Userid:    theUser.ID,
		Feedid:    dataParam.FeedId,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feedFollow: %v", err))
	}

	parsedFeedFollow := FeedFollow{
		ID:        feedFollow.ID,
		Createdat: feedFollow.Createdat,
		Updatedat: feedFollow.Updatedat,
		Feedid:    feedFollow.Feedid,
		Userid:    feedFollow.Userid,
	}

	respondWithJson(w, 201, parsedFeedFollow)
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, h *http.Request, theUser database.User) {

	feedFollowed, err := apiCfg.db.GetFeedFollow(h.Context(), theUser.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feedFollow: %v", err))
	}

	parsedFeedFollowed := []FeedFollow{}

	for _, feed := range feedFollowed {
		parsedFeedFollowed = append(parsedFeedFollowed, FeedFollow{
			ID:        feed.ID,
			Createdat: feed.Createdat,
			Updatedat: feed.Updatedat,
			Feedid:    feed.Feedid,
			Userid:    feed.Userid,
		})
	}

	respondWithJson(w, 201, parsedFeedFollowed)
}
func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, h *http.Request, theUser database.User) {

	feedFollowIdString := chi.URLParam(h, "feedFollowId")

	feedFollowId, err := uuid.Parse(feedFollowIdString)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("coudn't parse the string: %v", err))

	}

	deleteFeedFollowErr := apiCfg.db.DeleteFeedFollow(h.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		Userid: theUser.ID,
	})

	if deleteFeedFollowErr != nil {
		respondWithError(w, 403, fmt.Sprintf("couldn't delete feedFollow: %v", deleteFeedFollowErr))
	}

	respondWithJson(w, 200, "FeedFollow deleted")
}
