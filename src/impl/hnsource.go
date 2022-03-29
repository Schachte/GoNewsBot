package impl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"newssync/src"
	"os"

	"github.com/gocolly/colly"
)

type HackerNewsSource struct {
	Url          string
	Filemetadata string
}

func (s *HackerNewsSource) Scrape() (*src.Post, error) {
	var linksData []src.Post

	c := colly.NewCollector(
		colly.AllowedDomains(s.Url),
	)

	c.OnHTML(".athing", func(e *colly.HTMLElement) {
		link := e.ChildAttrs("a", "href")[1]
		title := e.ChildText("a.titlelink")

		lm := src.Post{
			Title:      title,
			Link:       link,
			SourceType: src.HACKER_NEWS_SOURCE,
		}

		linksData = append(linksData, lm)
	})

	location := fmt.Sprintf("https://%s", src.HACKER_NEWS_URL)
	c.Visit(location)
	return &linksData[0], nil
}

func (s *HackerNewsSource) GetPreviousUpload() src.History {
	f, err := os.OpenFile(s.Filemetadata, os.O_APPEND|os.O_CREATE|os.O_RDWR, 777)

	if err != nil {
		log.Panic(err)
	}

	bytes, err := ioutil.ReadAll(f)

	tmpUpload := src.History{}
	unmarshalErr := json.Unmarshal(bytes, &tmpUpload)

	if unmarshalErr != nil {
		log.Panic(unmarshalErr)
	}

	return tmpUpload
}

func (s *HackerNewsSource) WriteUploadMetadata(h *src.History) {
	f, _ := os.OpenFile(s.Filemetadata, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 777)

	historyJson, err := json.MarshalIndent(h, "", "  ")

	if err != nil {
		log.Panic(err)
	}

	f.Write(historyJson)
}
