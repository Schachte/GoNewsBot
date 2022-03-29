package main

import (
	"log"
	"newssync/boot"
	"newssync/src"
	"time"
)

func main() {
	boot.LoadEnvironment()
	boot.GenerateDatabases("databases", "source.db")
	reCheckSources()

	ticker := time.NewTicker(2 * time.Hour)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			log.Println("Rerunning Source Checker")
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
		_, alreadyExists := loadUpdatedSource(&sources[idx], &post)

		if !alreadyExists {
			uploadPostToSources(sources[idx], sinks, &post)
			log.Println("Article has been posted!")
			continue
		}
		log.Println("No update is required, this article was already posted")
	}

}

func uploadPostToSources(s src.Source, sinks []src.Sink, p *src.Post) {
	for _, sink := range sinks {
		uploaded, _ := sink.Upload(p)

		if uploaded {
			newHistory := src.History{
				LastUpdated:      time.Now().String(),
				LastArticleTitle: p.Title,
				Source:           p.SourceType,
			}

			s.WriteUpload(&newHistory)
		}
	}
}

func loadUpdatedSource(s *src.Source, post *src.Post) (src.Post, bool) {
	pastPost := (*s).CheckDuplicatePost(post)
	return *post, pastPost
}

func retrieveCurrentPostData(sources []src.Source) []src.Post {
	var currentPosts []src.Post

	for _, source := range sources {
		res, err := source.Scrape()

		if err != nil {
			log.Panic(err)
		}

		currentPosts = append(currentPosts, res)
	}

	return currentPosts
}
