package boot

import (
	"newssync/src"
	"newssync/src/impl"
)

// RetrieveSources will yield all configured sources we want to source data from to fan out to our sinks
func RetrieveSources() []src.Source {
	hackernews := &impl.HackerNewsSource{
		Url:          src.HACKER_NEWS_URL,
		Filemetadata: "metadata/hackernews.json",
	}

	return []src.Source{hackernews}
}

// RetrieveSinks retrieves all valid upload sinks to post new content to, such as Twitter or Reddit
func RetrieveSinks() []src.Sink {
	twitter := &impl.TwitterSink{
		Uri: "https://api.twitter.com/2/tweets",
		Auth: src.Credentials{
			// TODO: Reinstate the usage of these parameters
			ClientId:     "",
			ClientSecret: "",
			Extra:        "",
		},
	}

	return []src.Sink{twitter}
}
