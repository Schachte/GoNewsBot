package impl

import (
	"fmt"
	"log"
	"newssync/src"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterSink struct {
	Uri  string
	Auth src.Credentials
}

func (ts *TwitterSink) GetAuthentication() *src.Credentials {
	return &ts.Auth
}

func (ts *TwitterSink) Upload(p *src.Post) (bool, error) {
	config := oauth1.NewConfig(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	token := oauth1.NewToken(os.Getenv("OAUTH_ID"), os.Getenv("OAUTH_SECRET"))

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	tweetBody := fmt.Sprintf("[Ryan's Twitter Bot] - Check out this article from %s! %s\n%s", p.SourceType, p.Link, p.Title)
	_, _, err := client.Statuses.Update(tweetBody, nil)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return true, err
}
