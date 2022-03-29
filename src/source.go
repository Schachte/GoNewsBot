package src

type Post struct {
	Title      string
	Link       string
	SourceType string
}

type History struct {
	LastUpdated      string `json:"LastUpdated"`
	LastArticleTitle string `json:"LastArticleTitle"`
}

type Source interface {
	Scrape() (Post, error)
	GetPreviousUpload() History
	WriteUpload(*History)
}
