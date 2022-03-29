package main

import (
	"fmt"
	"log"
	"newssync/boot"
	"newssync/data"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.Println("Environment variables have been loaded correctly...")

	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			fmt.Println("Rerunning Source Checker")
			reCheckSources()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func reCheckSources() {
	sources := boot.RetrieveSources()
	sinks := boot.RetrieveSinks()
	posts := retrieveCurrentPostData(sources)

	for idx, post := range posts {
		_, _, requiresUpdate := loadUpdatedSource(sources[idx], post)

		if requiresUpdate {
			go uploadPostToSources(sources[idx], sinks, post)
			log.Println("Article has been posted!")
		} else {
			log.Println("No update is required, this article was already posted")
		}
	}

}

func uploadPostToSources(s data.Source, sinks []data.Sink, p *data.Post) {
	for _, sink := range sinks {
		uploaded, _ := sink.Upload(p)

		if uploaded {
			newHistory := data.History{
				LastUpdated:      time.Now().String(),
				LastArticleTitle: p.Title,
			}

			s.WriteUploadMetadata(&newHistory)
		}
	}
}

func loadUpdatedSource(s data.Source, post *data.Post) (*data.Post, string, bool) {
	pastPost := s.GetPreviousUpload()
	return post, pastPost.LastArticleTitle, post.Title != pastPost.LastArticleTitle
}

func retrieveCurrentPostData(sources []data.Source) []*data.Post {
	var currentPosts []*data.Post
	for _, source := range sources {
		res, err := source.Scrape()
		if err != nil {
			log.Panic(err)
		}
		currentPosts = append(currentPosts, res)
	}
	return currentPosts
}
