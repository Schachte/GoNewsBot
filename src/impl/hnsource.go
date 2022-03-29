package impl

import (
	"database/sql"
	"fmt"
	"log"
	"newssync/src"
	"newssync/src/util"

	"github.com/gocolly/colly"
	_ "github.com/mattn/go-sqlite3"
)

const DEFAULT_PROTOCOL = "https://"

type HackerNewsSource struct {
	Url          string
	Filemetadata string
	DatabaseLoc  string
}

func (s *HackerNewsSource) Scrape() (src.Post, error) {
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

	location := fmt.Sprintf("%s%s", DEFAULT_PROTOCOL, src.HACKER_NEWS_URL)
	c.Visit(location)
	return linksData[0], nil
}

func (s *HackerNewsSource) CheckDuplicatePost(p *src.Post) bool {
	sqliteDatabase, err := sql.Open("sqlite3", s.DatabaseLoc)
	defer sqliteDatabase.Close()

	if err != nil {
		log.Panic(err)
	}

	util.CreateTable(sqliteDatabase)

	return util.CheckIfPostHasBeenPosted(sqliteDatabase, p.Title)
}

func (s *HackerNewsSource) WriteUpload(h *src.History) {
	sqliteDatabase, err := sql.Open("sqlite3", s.DatabaseLoc)

	if err != nil {
		log.Panic(err)
	}

	writeErr := util.StoreNewPost(sqliteDatabase, h)

	if writeErr != nil {
		log.Panic(writeErr)
	}
}
