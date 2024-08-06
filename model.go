package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Userid    uuid.UUID `json:"userId"`
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	Createdat time.Time `json:"createdAt"`
	Updatedat time.Time `json:"updatedAt"`
	Feedid    uuid.UUID `json:"feedId"`
	Userid    uuid.UUID `json:"userId"`
}
