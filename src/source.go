package src

type Post struct {
	Title      string
	Link       string
	SourceType string
}

type History struct {
	LastUpdated      string `json:"LastUpdated"`
	LastArticleTitle string `json:"LastArticleTitle"`
	Source           string `json:"Source"`
}

type Source interface {
	Scrape() (Post, error)
	CheckDuplicatePost(*Post) bool
	WriteUpload(*History)
}
