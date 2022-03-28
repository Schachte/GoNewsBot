package main

import (
	"log"
	"newssync/boot"
	"newssync/data"
	"time"
)

func main() {
	sources := boot.RetrieveSources()
	sinks := boot.RetrieveSinks()

	currentPosts := retrieveCurrentPostData(sources)

	for idx, post := range currentPosts {
		_, _, requiresUpdate := loadUpdatedSource(sources[idx], post)

		if requiresUpdate {
			uploadPostToSources(sources[idx], sinks, post)
			log.Println("Article has been posted!")
		} else {
			log.Println("No update is required, this article was already posted")
		}
	}
}

func uploadPostToSources(s data.Source, sinks []data.Sink, p *data.Post) {
	for _, sink := range sinks {
		uploaded, _ := sink.Upload(p)

		newHistory := data.History{
			LastUpdated:      time.Now().String(),
			LastArticleTitle: p.Title,
		}

		if uploaded {
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
