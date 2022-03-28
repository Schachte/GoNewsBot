package boot

import (
	"newssync/data"
	"newssync/data/impl"
)

// RetrieveSources will yield all configured sources we want to source data from to fan out to our sinks
func RetrieveSources() []data.Source {
	hackernews := &impl.HackerNewsSource{
		Url:          data.HACKER_NEWS_URL,
		Filemetadata: "metadata/hackernews.json",
	}

	return []data.Source{hackernews}
}

// RetrieveSinks retrieves all valid upload sinks to post new content to, such as Twitter or Reddit
func RetrieveSinks() []data.Sink {
	twitter := &impl.TwitterSink{
		Uri: "https://api.twitter.com/2/tweets",
		Auth: data.Credentials{
			// TODO: Reinstate the usage of these parameters
			ClientId:     "",
			ClientSecret: "",
			Extra:        "",
		},
	}

	return []data.Sink{twitter}
}
