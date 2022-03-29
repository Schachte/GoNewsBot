package impl

import (
	"fmt"
	"newssync/data"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterSink struct {
	Uri  string
	Auth data.Credentials
}

func (ts *TwitterSink) GetAuthentication() *data.Credentials {
	return &ts.Auth
}

func (ts *TwitterSink) Upload(p *data.Post) (bool, error) {
	config := oauth1.NewConfig(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	token := oauth1.NewToken(os.Getenv("OAUTH_ID"), os.Getenv("OAUTH_SECRET"))

	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	tweetBody := fmt.Sprintf("[Ryan's Twitter Bot] - Check out this article from %s! %s\n%s", p.SourceType, p.Link, p.Title)
	_, _, err := client.Statuses.Update(tweetBody, nil)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, err
}

func (ts *TwitterSink) GetUri() string {
	return ts.Uri
}
