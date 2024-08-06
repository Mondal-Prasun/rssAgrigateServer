package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Mondal-Prasun/rssAgrigateServer/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Println("Scraping on ", concurrency, "corutiens every ", timeBetweenRequest, " duaration")

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error fetching feeds", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)

		}

		wg.Wait()

	}

}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched", err)
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Createdat:   time.Now().Local(),
			Updatedat:   time.Now().Local(),
			Title:       item.Title,
			Description: description,
			Publishedat: item.PubDate,
			Url:         item.Link,
			Feedid:      feed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Cant create post: ", err)
		}
	}

	log.Println("Feed ", feed.Name, "collected ", len(rssFeed.Channel.Item), "posts found")

}
